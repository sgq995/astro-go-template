import { spawn } from 'node:child_process';

/** @returns {import('astro').AstroIntegration} */
export default function air() {
  return {
    name: 'astro-air-integration',
    hooks: {
      'astro:server:start': async () => {
        const proc = spawn('air', {
          stdio: 'inherit',
          shell: true,
        });
        proc.on('error', console.error);
      },
      'astro:build:done': () => {
        const proc = spawn('go', ['run', 'cmd/build/main.go'], {
          stdio: 'inherit',
          shell: true,
        });
        proc.on('error', console.error);
      },
    },
  };
}
