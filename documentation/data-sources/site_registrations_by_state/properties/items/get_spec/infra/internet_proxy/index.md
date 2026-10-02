---
page_title: "items.get_spec.infra.internet_proxy"
subcategory: ""
description: "Proxy describes OPTIONS for HTTP or HTTPS proxy configurations."
xcsh_docs: {"aliases": ["items get spec infra internet proxy"], "body_bytes": 2447, "body_sha256": "sha256:70287da3da5a81b3b779d7e2ceb81a3f7ebebf4d9d3b16d6bfe0fcc40d2ae6cf", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:internet_proxy", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra", "path": "documentation/data-sources/site_registrations_by_state/properties/items/get_spec/infra/internet_proxy/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3123211222220010-3010203323202310-0310321132111233-2302131232321211-3310321023331323-1320100102101121-3030032120322001-1122300330012121", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "internet_proxy"], "schema_version": 1, "sections": [{"aliases": ["http proxy"], "anchor": "schema-items--get_spec--infra--internet_proxy--http_proxy", "description": "It will be used as the proxy URL for HTTP requests and HTTPS requests unless overridden by HTTPSProxy or NoProxy.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:internet_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "internet_proxy", "http_proxy"], "syntax": "attribute", "type": "string"}, {"aliases": ["https proxy"], "anchor": "schema-items--get_spec--infra--internet_proxy--https_proxy", "description": "It will be used as the proxy URL for HTTPS requests unless overridden by NoProxy.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:internet_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "internet_proxy", "https_proxy"], "syntax": "attribute", "type": "string"}, {"aliases": ["no proxy"], "anchor": "schema-items--get_spec--infra--internet_proxy--no_proxy", "description": "It specifies a string that contains comma-separated values specifying hosts that should be excluded from proxying. Each value is represented by an IP address prefix (192.0.2.103), an IP address prefix in CIDR notation (192.0.2.103/8), a domain name, or a special DNS label (*). An IP address..", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:internet_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "internet_proxy", "no_proxy"], "syntax": "attribute", "type": "string"}, {"aliases": ["cert", "certificate", "existing certificates", "proxy cacert url", "tls certificates"], "anchor": "schema-items--get_spec--infra--internet_proxy--proxy_cacert_url", "description": "Allow optional different trust-store for proxy in HTTP CONNECT step by picking proxy CA certificate value.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:infra:internet_proxy", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "internet_proxy", "proxy_cacert_url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/get_spec/infra/internet_proxy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Proxy describes OPTIONS for HTTP or HTTPS proxy configurations.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.internet_proxy

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/)
- items.get_spec.infra.internet_proxy

<a id="section"></a>

Type: `"single"`. Computed.

Proxy describes OPTIONS for HTTP or HTTPS proxy configurations.

## Direct properties

<a id="schema-items--get_spec--infra--internet_proxy--http_proxy"></a>

### http_proxy property

Type: `"string"`. Computed.

It will be used as the proxy URL for HTTP requests and HTTPS requests unless overridden by
HTTPSProxy or NoProxy.

<a id="schema-items--get_spec--infra--internet_proxy--https_proxy"></a>

### https_proxy property

Type: `"string"`. Computed.

It will be used as the proxy URL for HTTPS requests unless overridden by NoProxy.

<a id="schema-items--get_spec--infra--internet_proxy--no_proxy"></a>

### no_proxy property

Type: `"string"`. Computed.

It specifies a string that contains comma-separated values specifying hosts that should be excluded
from proxying. Each value is represented by an IP address prefix (192.0.2.103), an IP address prefix
in CIDR notation (192.0.2.103/8), a domain name, or a special DNS label (\*). An IP address..

<a id="schema-items--get_spec--infra--internet_proxy--proxy_cacert_url"></a>

### proxy_cacert_url property

Type: `"string"`. Computed.

Allow optional different trust-store for proxy in HTTP CONNECT step by picking proxy CA certificate
value.

## Next pages

- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/infra/)
- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
