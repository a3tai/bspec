function escapeHtml(value: string): string {
  return value
    .replaceAll('&', '&amp;')
    .replaceAll('<', '&lt;')
    .replaceAll('>', '&gt;')
    .replaceAll('"', '&quot;')
    .replaceAll("'", '&#39;');
}

function safeUrl(value: string): string {
  const trimmed = value.trim();
  if (/^(https?:|mailto:|#|\.{0,2}\/)/i.test(trimmed)) {
    return escapeHtml(trimmed);
  }
  return '#';
}

function renderInlineWithoutLinks(value: string): string {
  return escapeHtml(value)
    .replace(/`([^`]+)`/g, '<code>$1</code>')
    .replace(/\*\*([^*]+)\*\*/g, '<strong>$1</strong>')
    .replace(/__([^_]+)__/g, '<strong>$1</strong>')
    .replace(/\*([^*]+)\*/g, '<em>$1</em>')
    .replace(/_([^_]+)_/g, '<em>$1</em>');
}

function renderInline(value: string): string {
  const linkPattern = /\[([^\]]+)\]\(([^)\s]+)\)/g;
  let html = '';
  let cursor = 0;

  for (const match of value.matchAll(linkPattern)) {
    const index = match.index ?? 0;
    html += renderInlineWithoutLinks(value.slice(cursor, index));
    html += `<a href="${safeUrl(match[2])}" target="_blank" rel="noreferrer">${renderInlineWithoutLinks(match[1])}</a>`;
    cursor = index + match[0].length;
  }

  html += renderInlineWithoutLinks(value.slice(cursor));
  return html;
}

function splitTableRow(line: string): string[] {
  return line
    .trim()
    .replace(/^\|/, '')
    .replace(/\|$/, '')
    .split('|')
    .map((cell) => cell.trim());
}

function isTableSeparator(line: string): boolean {
  return /^\s*\|?\s*:?-{3,}:?\s*(\|\s*:?-{3,}:?\s*)+\|?\s*$/.test(line);
}

export function renderMarkdown(markdown: string): string {
  const source = markdown.trim();
  if (!source) {
    return '<p class="markdown-empty">No document body extracted.</p>';
  }

  const lines = source.replaceAll('\r\n', '\n').split('\n');
  const blocks: string[] = [];
  let paragraph: string[] = [];
  let listItems: string[] = [];
  let listTag: 'ul' | 'ol' | '' = '';
  let quote: string[] = [];
  let code: string[] = [];
  let codeLanguage = '';
  let inCode = false;

  function flushParagraph() {
    if (paragraph.length === 0) return;
    blocks.push(`<p>${renderInline(paragraph.join(' '))}</p>`);
    paragraph = [];
  }

  function flushList() {
    if (!listTag) return;
    blocks.push(`<${listTag}>${listItems.map((item) => `<li>${renderInline(item)}</li>`).join('')}</${listTag}>`);
    listItems = [];
    listTag = '';
  }

  function flushQuote() {
    if (quote.length === 0) return;
    blocks.push(`<blockquote>${renderMarkdown(quote.join('\n'))}</blockquote>`);
    quote = [];
  }

  function flushCode() {
    if (!inCode) return;
    const className = codeLanguage ? ` class="language-${escapeHtml(codeLanguage)}"` : '';
    blocks.push(`<pre><code${className}>${escapeHtml(code.join('\n'))}</code></pre>`);
    code = [];
    codeLanguage = '';
    inCode = false;
  }

  for (let index = 0; index < lines.length; index += 1) {
    const rawLine = lines[index];
    const line = rawLine.trimEnd();

    if (line.startsWith('```')) {
      if (inCode) {
        flushCode();
      } else {
        flushParagraph();
        flushList();
        flushQuote();
        inCode = true;
        codeLanguage = line.slice(3).trim().split(/\s+/)[0] ?? '';
      }
      continue;
    }

    if (inCode) {
      code.push(rawLine);
      continue;
    }

    if (line.trim() === '') {
      flushParagraph();
      flushList();
      flushQuote();
      continue;
    }

    if (index + 1 < lines.length && line.includes('|') && isTableSeparator(lines[index + 1])) {
      flushParagraph();
      flushList();
      flushQuote();

      const headers = splitTableRow(line);
      const rows: string[][] = [];
      index += 2;

      while (index < lines.length && lines[index].includes('|') && lines[index].trim() !== '') {
        rows.push(splitTableRow(lines[index]));
        index += 1;
      }
      index -= 1;

      blocks.push(
        `<table><thead><tr>${headers.map((cell) => `<th>${renderInline(cell)}</th>`).join('')}</tr></thead><tbody>${rows
          .map((row) => `<tr>${row.map((cell) => `<td>${renderInline(cell)}</td>`).join('')}</tr>`)
          .join('')}</tbody></table>`,
      );
      continue;
    }

    const heading = /^(#{1,6})\s+(.+)$/.exec(line);
    if (heading) {
      flushParagraph();
      flushList();
      flushQuote();
      const level = heading[1].length;
      blocks.push(`<h${level}>${renderInline(heading[2])}</h${level}>`);
      continue;
    }

    if (/^(-{3,}|\*{3,})\s*$/.test(line.trim())) {
      flushParagraph();
      flushList();
      flushQuote();
      blocks.push('<hr>');
      continue;
    }

    const unordered = /^\s*[-*+]\s+(.+)$/.exec(line);
    if (unordered) {
      flushParagraph();
      flushQuote();
      if (listTag && listTag !== 'ul') flushList();
      listTag = 'ul';
      listItems.push(unordered[1]);
      continue;
    }

    const ordered = /^\s*\d+\.\s+(.+)$/.exec(line);
    if (ordered) {
      flushParagraph();
      flushQuote();
      if (listTag && listTag !== 'ol') flushList();
      listTag = 'ol';
      listItems.push(ordered[1]);
      continue;
    }

    const quoted = /^\s*>\s?(.*)$/.exec(line);
    if (quoted) {
      flushParagraph();
      flushList();
      quote.push(quoted[1]);
      continue;
    }

    flushList();
    flushQuote();
    paragraph.push(line.trim());
  }

  flushCode();
  flushParagraph();
  flushList();
  flushQuote();

  return blocks.join('\n');
}
