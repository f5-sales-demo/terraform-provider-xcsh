---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_site_upgrade_status."
xcsh_docs: {"aliases": [], "body_bytes": 6318, "body_sha256": "sha256:30993f38b679a82c72f1cc8f73874ffc0169bbedfcf477f88abea8b2cf549010", "canonical_id": "xcsh-docs:data-sources:site_upgrade_status:reference", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_upgrade_status:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_upgrade_status:reference", "parent_id": "xcsh-docs:data-sources:site_upgrade_status:fundamentals", "path": "docs/guides/data-sources--site_upgrade_status--reference.md", "provider_name": "site_upgrade_status", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_upgrade_status/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_site_upgrade_status.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_site_upgrade_status](../data-sources/site_upgrade_status.md)
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
| `eligible` | [eligible](data-sources--site_upgrade_status--reference.md#schema-eligible) |
| `expected_os_version` | [expected_os_version](data-sources--site_upgrade_status--reference.md#schema-expected_os_version) |
| `expected_software_version` | [expected_software_version](data-sources--site_upgrade_status--reference.md#schema-expected_software_version) |
| `failed_precheck_names` | [failed_precheck_names](data-sources--site_upgrade_status--reference.md#schema-failed_precheck_names) |
| `id` | [id](data-sources--site_upgrade_status--reference.md#schema-id) |
| `os_available_version` | [os_available_version](data-sources--site_upgrade_status--reference.md#schema-os_available_version) |
| `os_deployment_phase` | [os_deployment_phase](data-sources--site_upgrade_status--reference.md#schema-os_deployment_phase) |
| `os_deployment_result` | [os_deployment_result](data-sources--site_upgrade_status--reference.md#schema-os_deployment_result) |
| `os_installed_version` | [os_installed_version](data-sources--site_upgrade_status--reference.md#schema-os_installed_version) |
| `poll_interval_seconds` | [poll_interval_seconds](data-sources--site_upgrade_status--reference.md#schema-poll_interval_seconds) |
| `ready` | [ready](data-sources--site_upgrade_status--reference.md#schema-ready) |
| `site` | [site](data-sources--site_upgrade_status--reference.md#schema-site) |
| `site_state` | [site_state](data-sources--site_upgrade_status--reference.md#schema-site_state) |
| `software_available_version` | [software_available_version](data-sources--site_upgrade_status--reference.md#schema-software_available_version) |
| `software_deployment_phase` | [software_deployment_phase](data-sources--site_upgrade_status--reference.md#schema-software_deployment_phase) |
| `software_deployment_result` | [software_deployment_result](data-sources--site_upgrade_status--reference.md#schema-software_deployment_result) |
| `software_installed_version` | [software_installed_version](data-sources--site_upgrade_status--reference.md#schema-software_installed_version) |
| `target_converged` | [target_converged](data-sources--site_upgrade_status--reference.md#schema-target_converged) |
| `timeout_seconds` | [timeout_seconds](data-sources--site_upgrade_status--reference.md#schema-timeout_seconds) |
| `upgradable_software_versions` | [upgradable_software_versions](data-sources--site_upgrade_status--reference.md#schema-upgradable_software_versions) |
| `wait` | [wait](data-sources--site_upgrade_status--reference.md#schema-wait) |

## Next pages

- [xcsh_site_upgrade_status](../data-sources/site_upgrade_status.md)
