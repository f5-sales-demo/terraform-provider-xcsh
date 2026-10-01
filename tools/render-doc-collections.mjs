// Render canonical collections independently; never bundle the whole corpus.
import { createHash } from 'node:crypto';
import { cp, mkdir, readFile, writeFile } from 'node:fs/promises';
import { dirname, join, resolve } from 'node:path';
import { createMarkdownProcessor } from '@astrojs/markdown-remark';

const [source, destination] = process.argv.slice(2);
if (!source || !destination) {
  throw new Error('usage: render-doc-collections.mjs <documentation> <output>');
}
const root = resolve(source);
const output = resolve(destination);
const manifest = JSON.parse(await readFile(join(root, 'generated-manifest.json'), 'utf8'));
const processor = await createMarkdownProcessor();
const escapeHtml = (value) => value.replaceAll('&', '&amp;').replaceAll('<', '&lt;').replaceAll('"', '&quot;');
const prefix = process.env.DOCUMENTATION_PREFIX || '';
const label = process.env.DOCUMENTATION_LABEL || 'Latest stable';
const versionBase = `/terraform-provider-xcsh/${prefix ? `${prefix}/` : ''}`;
let count = 0;
for (const [name, evidence] of Object.entries(manifest.files)) {
  if (!name.startsWith('documentation/') || !name.endsWith('.md') || name.includes('/_data/')) continue;
  const relative = name.slice('documentation/'.length);
  const bytes = await readFile(join(root, relative));
  if (bytes.length !== evidence.bytes || `sha256:${createHash('sha256').update(bytes).digest('hex')}` !== evidence.sha256) {
    throw new Error(`canonical page differs from manifest: ${relative}`);
  }
  const markdown = bytes.toString('utf8');
  const separator = markdown.indexOf('\n---\n\n', 4);
  if (!markdown.startsWith('---\n') || separator < 0) throw new Error(`invalid frontmatter: ${relative}`);
  const frontmatter = markdown.slice(4, separator);
  const titleLine = frontmatter.split('\n').find((line) => line.startsWith('page_title: '));
  const title = JSON.parse(titleLine.slice('page_title: '.length));
  const metadataLine = frontmatter.split('\n').find((line) => line.startsWith('xcsh_docs: '));
  const metadata = metadataLine ? JSON.parse(metadataLine.slice('xcsh_docs: '.length)) : null;
  const body = markdown.slice(separator + '\n---\n\n'.length);
  const { code } = await processor.render(body);
  const target = join(output, relative.replace(/\.md$/, '.html'));
  await mkdir(dirname(target), { recursive: true });
  const provenance = metadata ? `<footer>Schema path: <code>${escapeHtml(metadata.schema_path.join('.')) || 'root'}</code> · <a href="${escapeHtml(metadata.source_url || '')}">Complete Markdown</a><br><small>${escapeHtml(metadata.id)} · ${escapeHtml(metadata.body_sha256)}</small></footer>` : '';
  const html = `<!doctype html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"><title>${escapeHtml(title)} | xcsh</title><style>body{margin:0;color:#182b46;background:#f7f9fc;font:17px/1.65 system-ui,sans-serif}header{background:#102744;color:white;padding:1rem 2rem}header a{color:white}main{max-width:1100px;margin:2rem auto;padding:2rem;background:white;border-radius:12px}h1{font-size:2rem}h2{margin-top:2rem}a{color:#0758a8}pre{background:#eef2f6;padding:1rem;overflow:auto;border-radius:6px}code{font-size:.9em}table{border-collapse:collapse;width:100%;display:block;overflow:auto}td,th{border:1px solid #d6dfeb;padding:.5rem;text-align:left}footer{border-top:1px solid #d6dfeb;margin-top:2rem;padding-top:1rem;overflow-wrap:anywhere}small{color:#526078}</style></head><body><header><a href="${versionBase}">xcsh provider documentation</a> · ${escapeHtml(label)} · <a href="${versionBase}provider/setup/">Setup</a></header><main>${code}${provenance}</main></body></html>\n`;
  await writeFile(target, html);
  count++;
}
await cp(join(root, '_data'), join(output, '_data'), { recursive: true });
for (const file of ['llms.txt', 'terraform-llms-index.json', 'generated-manifest.json']) {
  await cp(join(root, file), join(output, file));
}
console.log(`Rendered ${count} complete canonical pages independently.`);

const search = Object.keys(manifest.files).filter((name) => name.startsWith('documentation/') && name.endsWith('.md') && !name.includes('/_data/')).map((name) => ({ path: versionBase + name.slice('documentation/'.length).replace(/index\.md$/, '').replace(/\.md$/, '/'), title: name.slice('documentation/'.length) }));
await writeFile(join(output, 'search-index.json'), JSON.stringify(search) + '\n');
