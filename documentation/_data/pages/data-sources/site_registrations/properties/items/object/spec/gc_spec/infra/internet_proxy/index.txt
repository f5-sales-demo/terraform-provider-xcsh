---
page_title: "items.object.spec.gc_spec.infra.internet_proxy"
subcategory: ""
description: "Proxy describes OPTIONS for HTTP or HTTPS proxy configurations."
xcsh_docs: {"aliases": ["items object spec gc spec infra internet proxy"], "body_bytes": 2778, "body_sha256": "sha256:f3f2a5bbe502b1a7a8ca31c1191a7a78325a763d880a893c6a4f817802d42310", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:internet_proxy", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra", "path": "documentation/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/internet_proxy/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0231310112131331-0121003031132331-1010302013100213-3323231013331231-1103220230130021-0300201300303201-3033311202130313-0212323313221301", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra", "internet_proxy"], "schema_version": 1, "sections": [{"aliases": ["http proxy"], "anchor": "schema-items--object--spec--gc_spec--infra--internet_proxy--http_proxy", "description": "It will be used as the proxy URL for HTTP requests and HTTPS requests unless overridden by HTTPSProxy or NoProxy.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:internet_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "internet_proxy", "http_proxy"], "syntax": "attribute", "type": "string"}, {"aliases": ["https proxy"], "anchor": "schema-items--object--spec--gc_spec--infra--internet_proxy--https_proxy", "description": "It will be used as the proxy URL for HTTPS requests unless overridden by NoProxy.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:internet_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "internet_proxy", "https_proxy"], "syntax": "attribute", "type": "string"}, {"aliases": ["no proxy"], "anchor": "schema-items--object--spec--gc_spec--infra--internet_proxy--no_proxy", "description": "It specifies a string that contains comma-separated values specifying hosts that should be excluded from proxying. Each value is represented by an IP address prefix (192.0.2.103), an IP address prefix in CIDR notation (192.0.2.103/8), a domain name, or a special DNS label (*). An IP address..", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:internet_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "internet_proxy", "no_proxy"], "syntax": "attribute", "type": "string"}, {"aliases": ["cert", "certificate", "existing certificates", "proxy cacert url", "tls certificates"], "anchor": "schema-items--object--spec--gc_spec--infra--internet_proxy--proxy_cacert_url", "description": "Allow optional different trust-store for proxy in HTTP CONNECT step by picking proxy CA certificate value.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:object:spec:gc_spec:infra:internet_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "spec", "gc_spec", "infra", "internet_proxy", "proxy_cacert_url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/internet_proxy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Proxy describes OPTIONS for HTTP or HTTPS proxy configurations.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.infra.internet_proxy

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/)
- [items.object.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/)
- [items.object.spec.gc_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/)
- [items.object.spec.gc_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/)
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

- [items.object.spec.gc_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/object/spec/gc_spec/infra/)
- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
