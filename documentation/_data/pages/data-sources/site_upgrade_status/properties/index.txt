---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_site_upgrade_status."
xcsh_docs: {"aliases": [], "body_bytes": 7434, "body_sha256": "sha256:3fd9b9bf9a139f57d2267db77be27c29f3d92b971d9c04a378c4a2ece2cca8d7", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_upgrade_status:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_upgrade_status:reference", "parent_id": "xcsh-docs:data-sources:site_upgrade_status:fundamentals", "path": "documentation/data-sources/site_upgrade_status/properties/index.md", "provider_name": "site_upgrade_status", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_upgrade_status/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_site_upgrade_status.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_site_upgrade_status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/)
- Property reference

## Direct properties

<a id="schema-eligible"></a>

### eligible property

Type: `"bool"`. Computed.

Whether the site is ONLINE and each selected target is installed or advertised for upgrade. Software
prechecks must pass when software would change; an unchanged paired version does not block a serial
software or OS upgrade.

<a id="schema-expected_os_version"></a>

### expected_os_version property

Type: `"string"`. Optional.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="schema-expected_software_version"></a>

### expected_software_version property

Type: `"string"`. Optional.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="schema-failed_precheck_names"></a>

### failed_precheck_names property

Type: `["list", "string"]`. Computed.

Failed software prechecks for a newer software target; empty when the selected software version is
already installed.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

<a id="schema-os_available_version"></a>

### os_available_version property

Type: `"string"`. Computed.

<a id="schema-os_deployment_phase"></a>

### os_deployment_phase property

Type: `"string"`. Computed.

<a id="schema-os_deployment_result"></a>

### os_deployment_result property

Type: `"string"`. Computed.

<a id="schema-os_installed_version"></a>

### os_installed_version property

Type: `"string"`. Computed.

<a id="schema-poll_interval_seconds"></a>

### poll_interval_seconds property

Type: `"number"`. Optional, Computed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{int64validator.Between(1, 300)}
```

<a id="schema-ready"></a>

### ready property

Type: `"bool"`. Computed.

Whether the site is operationally ready (\`ONLINE\`), independent of target eligibility.

<a id="schema-site"></a>

### site property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="schema-site_state"></a>

### site_state property

Type: `"string"`. Computed.

<a id="schema-software_available_version"></a>

### software_available_version property

Type: `"string"`. Computed.

<a id="schema-software_deployment_phase"></a>

### software_deployment_phase property

Type: `"string"`. Computed.

<a id="schema-software_deployment_result"></a>

### software_deployment_result property

Type: `"string"`. Computed.

<a id="schema-software_installed_version"></a>

### software_installed_version property

Type: `"string"`. Computed.

<a id="schema-target_converged"></a>

### target_converged property

Type: `"bool"`. Computed.

<a id="schema-timeout_seconds"></a>

### timeout_seconds property

Type: `"number"`. Optional, Computed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{int64validator.Between(1, 7200)}
```

<a id="schema-upgradable_software_versions"></a>

### upgradable_software_versions property

Type: `["list", "string"]`. Computed.

<a id="schema-wait"></a>

### wait property

Type: `"bool"`. Optional, Computed.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `eligible` | [eligible](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/properties/#schema-eligible) |
| `expected_os_version` | [expected_os_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/properties/#schema-expected_os_version) |
| `expected_software_version` | [expected_software_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/properties/#schema-expected_software_version) |
| `failed_precheck_names` | [failed_precheck_names](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/properties/#schema-failed_precheck_names) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/properties/#schema-id) |
| `os_available_version` | [os_available_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/properties/#schema-os_available_version) |
| `os_deployment_phase` | [os_deployment_phase](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/properties/#schema-os_deployment_phase) |
| `os_deployment_result` | [os_deployment_result](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/properties/#schema-os_deployment_result) |
| `os_installed_version` | [os_installed_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/properties/#schema-os_installed_version) |
| `poll_interval_seconds` | [poll_interval_seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/properties/#schema-poll_interval_seconds) |
| `ready` | [ready](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/properties/#schema-ready) |
| `site` | [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/properties/#schema-site) |
| `site_state` | [site_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/properties/#schema-site_state) |
| `software_available_version` | [software_available_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/properties/#schema-software_available_version) |
| `software_deployment_phase` | [software_deployment_phase](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/properties/#schema-software_deployment_phase) |
| `software_deployment_result` | [software_deployment_result](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/properties/#schema-software_deployment_result) |
| `software_installed_version` | [software_installed_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/properties/#schema-software_installed_version) |
| `target_converged` | [target_converged](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/properties/#schema-target_converged) |
| `timeout_seconds` | [timeout_seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/properties/#schema-timeout_seconds) |
| `upgradable_software_versions` | [upgradable_software_versions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/properties/#schema-upgradable_software_versions) |
| `wait` | [wait](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/properties/#schema-wait) |

## Next pages

- [xcsh_site_upgrade_status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/)
