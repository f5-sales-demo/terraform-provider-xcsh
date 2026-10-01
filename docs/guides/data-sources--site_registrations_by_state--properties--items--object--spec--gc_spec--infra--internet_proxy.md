---
page_title: "items.object.spec.gc_spec.infra.internet_proxy"
subcategory: ""
description: "items.object.spec.gc_spec.infra.internet_proxy for xcsh_site_registrations_by_state."
xcsh_docs: {"aliases": [], "body_bytes": 2430, "body_sha256": "sha256:f5e3be7191032b273dde95b22fefb672b885e0eef521e4fce23be704bc3e0cc6", "canonical_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:internet_proxy", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra:internet_proxy", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:infra", "path": "docs/guides/data-sources--site_registrations_by_state--properties--items--object--spec--gc_spec--infra--internet_proxy.md", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "infra", "internet_proxy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/infra/internet_proxy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.object.spec.gc_spec.infra.internet_proxy for xcsh_site_registrations_by_state.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.infra.internet_proxy

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md)
- [Property reference](data-sources--site_registrations_by_state--reference.md)
- [items](data-sources--site_registrations_by_state--properties--items.md)
- [items.object](data-sources--site_registrations_by_state--properties--items--object.md)
- [items.object.spec](data-sources--site_registrations_by_state--properties--items--object--spec.md)
- [items.object.spec.gc_spec](data-sources--site_registrations_by_state--properties--items--object--spec--gc_spec.md)
- [items.object.spec.gc_spec.infra](data-sources--site_registrations_by_state--properties--items--object--spec--gc_spec--infra.md)
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

- [items.object.spec.gc_spec.infra](data-sources--site_registrations_by_state--properties--items--object--spec--gc_spec--infra.md)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md)
