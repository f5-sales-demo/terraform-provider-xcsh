---
page_title: "items.object.spec.gc_spec.passport"
subcategory: ""
description: "items.object.spec.gc_spec.passport for xcsh_site_registrations_by_state."
xcsh_docs: {"aliases": [], "body_bytes": 4334, "body_sha256": "sha256:0314ea6e8327fc43d69db52b9ad9a22341be4ad4d1994b2b2486e4ef92d507d5", "canonical_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:passport", "child_ids": ["xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:passport:default_os_version", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:passport:default_sw_version"], "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec:passport", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:spec:gc_spec", "path": "docs/guides/data-sources--site_registrations_by_state--properties--items--object--spec--gc_spec--passport.md", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["items", "object", "spec", "gc_spec", "passport"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/object/spec/gc_spec/passport/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.object.spec.gc_spec.passport for xcsh_site_registrations_by_state.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.spec.gc_spec.passport

Breadcrumbs:

- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md)
- [Property reference](data-sources--site_registrations_by_state--reference.md)
- [items](data-sources--site_registrations_by_state--properties--items.md)
- [items.object](data-sources--site_registrations_by_state--properties--items--object.md)
- [items.object.spec](data-sources--site_registrations_by_state--properties--items--object--spec.md)
- [items.object.spec.gc_spec](data-sources--site_registrations_by_state--properties--items--object--spec--gc_spec.md)
- items.object.spec.gc_spec.passport

<a id="section"></a>

Type: `"single"`. Computed.

Passport stores information about identification and node configuration provided by CE during
registration. It can be manually updated by user during approval.

## Direct properties

<a id="schema-items--object--spec--gc_spec--passport--cluster_name"></a>

### cluster_name property

Type: `"string"`. Computed.

Cluster Name. Human-readable name for the resource

<a id="schema-items--object--spec--gc_spec--passport--cluster_size"></a>

### cluster_size property

Type: `"number"`. Computed.

Defines how many master nodes is in the cluster, only 1 or 3 is allowed 1 - cluster have single
master, without HA 3 - cluster have 3 masters, with HA, all nodes should be allowed at same time,
cluster won't start until ALL nodes are ADMITTED 0 - same as 1 This value can't be changed after..

<a id="schema-items--object--spec--gc_spec--passport--cluster_type"></a>

### cluster_type property

Type: `"string"`. Computed.

Cluster Type. Cluster or grouping configuration

- [default_os_version](data-sources--site_registrations_by_state--properties--items--object--spec--gc_spec--passport--default_os_version.md): complete subsection reference.

- [default_sw_version](data-sources--site_registrations_by_state--properties--items--object--spec--gc_spec--passport--default_sw_version.md): complete subsection reference.

<a id="schema-items--object--spec--gc_spec--passport--latitude"></a>

### latitude property

Type: `"number"`. Computed.

Latitude. Geographic location of this site.

<a id="schema-items--object--spec--gc_spec--passport--longitude"></a>

### longitude property

Type: `"number"`. Computed.

Longitude. Geographic location of this site.

<a id="schema-items--object--spec--gc_spec--passport--operating_system_version"></a>

### operating_system_version property

Type: `"string"`. Computed.

Exclusive with \[default\_os\_version\] Operating System Version is optional parameter, which allows
to specify target SW version for particular site e.g. 7.2009.10.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

<a id="schema-items--object--spec--gc_spec--passport--private_network_name"></a>

### private_network_name property

Type: `"string"`. Computed.

Private Network name for private access connectivity to F5XC ADN. It is used for PrivateLink,
CloudLink and L3VPN.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(64),
}
```

<a id="schema-items--object--spec--gc_spec--passport--volterra_software_version"></a>

### volterra_software_version property

Type: `"string"`. Computed.

Exclusive with \[default\_sw\_version\] F5XC Software Version is optional parameter, which allows to
specify target SW version for particular site e.g. Crt-20210329-1002.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(20),
}
```

## Next pages

- [items.object.spec.gc_spec.passport.default_os_version](data-sources--site_registrations_by_state--properties--items--object--spec--gc_spec--passport--default_os_version.md)
- [items.object.spec.gc_spec.passport.default_sw_version](data-sources--site_registrations_by_state--properties--items--object--spec--gc_spec--passport--default_sw_version.md)
- [items.object.spec.gc_spec](data-sources--site_registrations_by_state--properties--items--object--spec--gc_spec.md)
- [xcsh_site_registrations_by_state](../data-sources/site_registrations_by_state.md)
