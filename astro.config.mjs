// @ts-check
import { defineConfig } from 'astro/config';
import air from './integrations/air.mjs';

// https://astro.build/config
export default defineConfig({
  integrations: [air()]
});
