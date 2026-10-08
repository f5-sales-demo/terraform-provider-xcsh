---
page_title: "items.get_spec.infra.internet_proxy"
subcategory: ""
description: "Proxy describes OPTIONS for HTTP or HTTPS proxy configurations."
xcsh_docs: {"aliases": ["items get spec infra internet proxy"], "body_bytes": 2133, "body_sha256": "sha256:75a834377179ca7051a57a170f5b275e76b2abd7075427fbe60b94ac3e04933b", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:internet_proxy", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra", "path": "documentation/data-sources/site_registrations_by_site/properties/items/get_spec/infra/internet_proxy/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0202001222201213-0010202212131210-2110012030331330-3303110200113301-0111001121232300-3322103223301233-1002123202120301-0002132333222120", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "get_spec", "infra", "internet_proxy"], "schema_version": 1, "sections": [{"aliases": ["items get spec infra internet proxy http proxy"], "anchor": "schema-items--get_spec--infra--internet_proxy--http_proxy", "description": "It will be used as the proxy URL for HTTP requests and HTTPS requests unless overridden by HTTPSProxy or NoProxy.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:internet_proxy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "internet_proxy", "http_proxy"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra internet proxy https proxy"], "anchor": "schema-items--get_spec--infra--internet_proxy--https_proxy", "description": "It will be used as the proxy URL for HTTPS requests unless overridden by NoProxy.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:internet_proxy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "internet_proxy", "https_proxy"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra internet proxy no proxy"], "anchor": "schema-items--get_spec--infra--internet_proxy--no_proxy", "description": "It specifies a string that contains comma-separated values specifying hosts that should be excluded from proxying. Each value is represented by an IP address prefix (192.0.2.103), an IP address prefix in CIDR notation (192.0.2.103/8), a domain name, or a special DNS label (*). An IP address..", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:internet_proxy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "internet_proxy", "no_proxy"], "syntax": "attribute", "type": "string"}, {"aliases": ["items get spec infra internet proxy proxy cacert url"], "anchor": "schema-items--get_spec--infra--internet_proxy--proxy_cacert_url", "description": "Allow optional different trust-store for proxy in HTTP CONNECT step by picking proxy CA certificate value.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:get_spec:infra:internet_proxy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "get_spec", "infra", "internet_proxy", "proxy_cacert_url"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/get_spec/infra/internet_proxy/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Proxy describes OPTIONS for HTTP or HTTPS proxy configurations.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.infra.internet_proxy

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/)
- [items.get_spec.infra](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/get_spec/infra/)
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
