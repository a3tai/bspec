import { describe, expect, test } from 'bun:test';
import { buildGraphScene, serviceError } from '../src/lib/business';

describe('serviceError', () => {
  test('separates stable service codes from readable messages', () => {
    expect(serviceError(new Error('BUSINESS_PRECONDITION_FAILED: record version is stale'))).toEqual({
      code: 'BUSINESS_PRECONDITION_FAILED',
      message: 'record version is stale',
    });
  });
});

describe('buildGraphScene', () => {
  test('keeps unresolved references visible as terminal nodes', () => {
    const scene = buildGraphScene(
      {
        business_id: 'business',
        root_record_id: 'record',
        depth: 1,
        nodes: [
          {
            id: 'record',
            key: 'strategy',
            type_code: 'STR',
            title: 'Strategy',
            status: 'active',
            visibility: 'internal',
            version: 1,
          },
        ],
        edges: [
          {
            id: 'edge',
            source_id: 'record',
            target_ref: 'missing-plan',
            kind: 'depends_on',
            strength: '',
          },
        ],
        projected_through: '',
        truncated: false,
      } as never,
      [],
    );

    expect(scene.nodes).toHaveLength(2);
    expect(scene.nodes.find((node) => node.unresolved)?.key).toBe('missing-plan');
    expect(scene.edges[0].unresolved).toBe(true);
  });

  test('keeps dense nodes separate and routes parallel facts independently', () => {
    const nodes = Array.from({ length: 20 }, (_, index) => ({
      id: `record-${index}`,
      key: `record-${index}`,
      type_code: 'REC',
      title: `Record ${index}`,
      status: 'active',
      visibility: 'internal',
      version: 1,
    }));
    const scene = buildGraphScene(
      {
        business_id: 'business',
        root_record_id: 'record-0',
        depth: 2,
        nodes,
        edges: [
          { id: 'forward', source_id: 'record-0', target_id: 'record-1', kind: 'enables', strength: '' },
          { id: 'reverse', source_id: 'record-1', target_id: 'record-0', kind: 'related', strength: '' },
        ],
        projected_through: '',
        truncated: false,
      } as never,
      [],
    );

    for (let left = 0; left < scene.nodes.length; left += 1) {
      for (let right = left + 1; right < scene.nodes.length; right += 1) {
        const horizontal = Math.abs(scene.nodes[left].x - scene.nodes[right].x);
        const vertical = Math.abs(scene.nodes[left].y - scene.nodes[right].y);
        expect(horizontal >= 108 || vertical >= 54).toBe(true);
      }
    }
    expect(scene.edges.every((edge) => edge.path.includes(' Q '))).toBe(true);
    expect(new Set(scene.edges.map((edge) => `${edge.labelX}:${edge.labelY}`)).size).toBe(2);
  });
});
