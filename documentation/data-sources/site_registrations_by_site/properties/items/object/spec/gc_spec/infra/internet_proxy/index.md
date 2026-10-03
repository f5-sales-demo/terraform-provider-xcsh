---
page_title: "items.object.spec.gc_spec.infra.internet_proxy"
subcategory: ""
description: "Proxy describes OPTIONS for HTTP or HTTPS proxy configurations."
xcsh_docs: {"aliases": ["items object spec gc spec infra internet proxy"], "body_bytes": 2866, "body_sha256": "sha256:c039ba7884a8cc6f2670af1ce9a535017c736bc55c20bbc7195703c7f7248bdb", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:internet_proxy", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra", "path": "documentation/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/internet_proxy/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1121232113322103-1303112320303303-0113101133220101-1012312220121032-1110131113030213-3301003113333313-1113031203002000-0111310111320012", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra", "internet_proxy"], "schema_version": 1, "sections": [{"aliases": ["items object spec gc spec infra internet proxy http proxy"], "anchor": "schema-items--object--spec--gc_spec--infra--internet_proxy--http_proxy", "description": "It will be used as the proxy URL for HTTP requests and HTTPS requests unless overridden by HTTPSProxy or NoProxy.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:internet_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "internet_proxy", "http_proxy"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec infra internet proxy https proxy"], "anchor": "schema-items--object--spec--gc_spec--infra--internet_proxy--https_proxy", "description": "It will be used as the proxy URL for HTTPS requests unless overridden by NoProxy.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:internet_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "internet_proxy", "https_proxy"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec infra internet proxy no proxy"], "anchor": "schema-items--object--spec--gc_spec--infra--internet_proxy--no_proxy", "description": "It specifies a string that contains comma-separated values specifying hosts that should be excluded from proxying. Each value is represented by an IP address prefix (192.0.2.103), an IP address prefix in CIDR notation (192.0.2.103/8), a domain name, or a special DNS label (*). An IP address..", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:internet_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "internet_proxy", "no_proxy"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object spec gc spec infra internet proxy proxy cacert url"], "anchor": "schema-items--object--spec--gc_spec--infra--internet_proxy--proxy_cacert_url", "description": "Allow optional different trust-store for proxy in HTTP CONNECT step by picking proxy CA certificate value.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:spec:gc_spec:infra:internet_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "internet_proxy", "proxy_cacert_url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/internet_proxy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Proxy describes OPTIONS for HTTP or HTTPS proxy configurations.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.infra.internet_proxy

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/)
- [items.object.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/)
- [items.object.spec.gc_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/)
- [items.object.spec.gc_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/)
- items.object.spec.gc_spec.infra.internet_proxy

<a id="section"></a>

Type: `"single"`. Computed.

Proxy describes OPTIONS for HTTP or HTTPS proxy configurations.

## Direct properties

<a id="schema-items--object--spec--gc_spec--infra--internet_proxy--http_proxy"></a>

### http_proxy property

Type: `"string"`. Computed.

It will be used as the proxy URL for HTTP requests and HTTPS requests unless overridden by
HTTPSProxy or NoProxy.

<a id="schema-items--object--spec--gc_spec--infra--internet_proxy--https_proxy"></a>

### https_proxy property

Type: `"string"`. Computed.

It will be used as the proxy URL for HTTPS requests unless overridden by NoProxy.

<a id="schema-items--object--spec--gc_spec--infra--internet_proxy--no_proxy"></a>

### no_proxy property

Type: `"string"`. Computed.

It specifies a string that contains comma-separated values specifying hosts that should be excluded
from proxying. Each value is represented by an IP address prefix (192.0.2.103), an IP address prefix
in CIDR notation (192.0.2.103/8), a domain name, or a special DNS label (\*). An IP address..

<a id="schema-items--object--spec--gc_spec--infra--internet_proxy--proxy_cacert_url"></a>

### proxy_cacert_url property

Type: `"string"`. Computed.

Allow optional different trust-store for proxy in HTTP CONNECT step by picking proxy CA certificate
value.

## Next pages

- [items.object.spec.gc_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/spec/gc_spec/infra/)
- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
