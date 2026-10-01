---
page_title: "items.get_spec.passport"
subcategory: ""
description: "items.get_spec.passport for xcsh_site_registrations_by_state."
xcsh_docs: {"aliases": [], "body_bytes": 4393, "body_sha256": "sha256:1f3c21bbf3c0c6af22e097623393574fbc218f8a5a8ae2baeb3f7d85f89c87f6", "child_ids": ["xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:passport:default_os_version", "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:passport:default_sw_version"], "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec:passport", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:get_spec", "path": "documentation/data-sources/site_registrations_by_state/properties/items/get_spec/passport/index.md", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["items", "get_spec", "passport"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/get_spec/passport/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.get_spec.passport for xcsh_site_registrations_by_state.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.get_spec.passport

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/)
- items.get_spec.passport

<a id="section"></a>

Type: `"single"`. Computed.

Passport stores information about identification and node configuration provided by CE during
registration. It can be manually updated by user during approval.

## Direct properties

<a id="schema-items--get_spec--passport--cluster_name"></a>

### cluster_name property

Type: `"string"`. Computed.

Cluster Name. Human-readable name for the resource

<a id="schema-items--get_spec--passport--cluster_size"></a>

### cluster_size property

Type: `"number"`. Computed.

Defines how many master nodes is in the cluster, only 1 or 3 is allowed 1 - cluster have single
master, without HA 3 - cluster have 3 masters, with HA, all nodes should be allowed at same time,
cluster won't start until ALL nodes are ADMITTED 0 - same as 1 This value can't be changed after..

<a id="schema-items--get_spec--passport--cluster_type"></a>

### cluster_type property

Type: `"string"`. Computed.

Cluster Type. Cluster or grouping configuration

- [default_os_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/passport/default_os_version/): complete subsection reference.

- [default_sw_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/passport/default_sw_version/): complete subsection reference.

<a id="schema-items--get_spec--passport--latitude"></a>

### latitude property

Type: `"number"`. Computed.

Latitude. Geographic location of this site.

<a id="schema-items--get_spec--passport--longitude"></a>

### longitude property

Type: `"number"`. Computed.

Longitude. Geographic location of this site.

<a id="schema-items--get_spec--passport--operating_system_version"></a>

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

<a id="schema-items--get_spec--passport--private_network_name"></a>

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

<a id="schema-items--get_spec--passport--volterra_software_version"></a>

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

- [items.get_spec.passport.default_os_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/passport/default_os_version/)
- [items.get_spec.passport.default_sw_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/passport/default_sw_version/)
- [items.get_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/get_spec/)
- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
