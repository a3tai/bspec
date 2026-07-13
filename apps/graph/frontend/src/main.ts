import './styles/tokens.css';
import './styles/ui.css';
import './styles/app.css';

import { mount } from 'svelte';
import App from './App.svelte';

mount(App, { target: document.getElementById('app')! });
