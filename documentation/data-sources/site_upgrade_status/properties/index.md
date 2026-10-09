---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_site_upgrade_status."
xcsh_docs: {"aliases": ["site upgrade status"], "body_bytes": 7544, "body_sha256": "sha256:4ad3e8a2bd49161f2950fe066df580398250bd3b4ce707991803439060a92f4d", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_upgrade_status:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_upgrade_status:reference", "parent_id": "xcsh-docs:data-sources:site_upgrade_status:fundamentals", "path": "documentation/data-sources/site_upgrade_status/properties/index.md", "product": "distributed-cloud", "provider_name": "site_upgrade_status", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0002123201223230-3102222113003200-2112313021101312-3221113322221231-2011313330023232-1101331323313032-0200231210030103-1020333122010212", "registry_path": "docs/guides/data-sources--site_upgrade_status--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["eligible"], "anchor": "schema-eligible", "description": "Whether the site is ONLINE and each selected target is installed or advertised for upgrade. Software prechecks must pass when software would change; an unchanged paired version does not block a serial software or OS upgrade.", "document_id": "xcsh-docs:data-sources:site_upgrade_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["eligible"], "syntax": "attribute", "type": "bool"}, {"aliases": ["expected os version"], "anchor": "schema-expected_os_version", "description": "expected os version", "document_id": "xcsh-docs:data-sources:site_upgrade_status:reference", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["expected_os_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["expected software version"], "anchor": "schema-expected_software_version", "description": "expected software version", "document_id": "xcsh-docs:data-sources:site_upgrade_status:reference", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["expected_software_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["failed precheck names"], "anchor": "schema-failed_precheck_names", "description": "Failed software prechecks for a newer software target; empty when the selected software version is already installed.", "document_id": "xcsh-docs:data-sources:site_upgrade_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["failed_precheck_names"], "syntax": "attribute", "type": "list"}, {"aliases": ["id"], "anchor": "schema-id", "description": "id", "document_id": "xcsh-docs:data-sources:site_upgrade_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["os available version"], "anchor": "schema-os_available_version", "description": "os available version", "document_id": "xcsh-docs:data-sources:site_upgrade_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["os_available_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["os deployment phase"], "anchor": "schema-os_deployment_phase", "description": "os deployment phase", "document_id": "xcsh-docs:data-sources:site_upgrade_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["os_deployment_phase"], "syntax": "attribute", "type": "string"}, {"aliases": ["os deployment result"], "anchor": "schema-os_deployment_result", "description": "os deployment result", "document_id": "xcsh-docs:data-sources:site_upgrade_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["os_deployment_result"], "syntax": "attribute", "type": "string"}, {"aliases": ["os installed version"], "anchor": "schema-os_installed_version", "description": "os installed version", "document_id": "xcsh-docs:data-sources:site_upgrade_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["os_installed_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["poll interval seconds"], "anchor": "schema-poll_interval_seconds", "description": "poll interval seconds", "document_id": "xcsh-docs:data-sources:site_upgrade_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["poll_interval_seconds"], "syntax": "attribute", "type": "number"}, {"aliases": ["ready"], "anchor": "schema-ready", "description": "Whether the site is operationally ready (`ONLINE`), independent of target eligibility.", "document_id": "xcsh-docs:data-sources:site_upgrade_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["ready"], "syntax": "attribute", "type": "bool"}, {"aliases": ["site"], "anchor": "schema-site", "description": "site", "document_id": "xcsh-docs:data-sources:site_upgrade_status:reference", "enum_extraction_complete": true, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site"], "syntax": "attribute", "type": "string"}, {"aliases": ["site state"], "anchor": "schema-site_state", "description": "site state", "document_id": "xcsh-docs:data-sources:site_upgrade_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_state"], "syntax": "attribute", "type": "string"}, {"aliases": ["software available version"], "anchor": "schema-software_available_version", "description": "software available version", "document_id": "xcsh-docs:data-sources:site_upgrade_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["software_available_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["software deployment phase"], "anchor": "schema-software_deployment_phase", "description": "software deployment phase", "document_id": "xcsh-docs:data-sources:site_upgrade_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["software_deployment_phase"], "syntax": "attribute", "type": "string"}, {"aliases": ["software deployment result"], "anchor": "schema-software_deployment_result", "description": "software deployment result", "document_id": "xcsh-docs:data-sources:site_upgrade_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["software_deployment_result"], "syntax": "attribute", "type": "string"}, {"aliases": ["software installed version"], "anchor": "schema-software_installed_version", "description": "software installed version", "document_id": "xcsh-docs:data-sources:site_upgrade_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["software_installed_version"], "syntax": "attribute", "type": "string"}, {"aliases": ["target converged"], "anchor": "schema-target_converged", "description": "target converged", "document_id": "xcsh-docs:data-sources:site_upgrade_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["target_converged"], "syntax": "attribute", "type": "bool"}, {"aliases": ["duration", "timeout seconds"], "anchor": "schema-timeout_seconds", "description": "timeout seconds", "document_id": "xcsh-docs:data-sources:site_upgrade_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["timeout_seconds"], "syntax": "attribute", "type": "number"}, {"aliases": ["upgradable software versions"], "anchor": "schema-upgradable_software_versions", "description": "upgradable software versions", "document_id": "xcsh-docs:data-sources:site_upgrade_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["upgradable_software_versions"], "syntax": "attribute", "type": "list"}, {"aliases": ["wait"], "anchor": "schema-wait", "description": "wait", "document_id": "xcsh-docs:data-sources:site_upgrade_status:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["wait"], "syntax": "attribute", "type": "bool"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_upgrade_status/properties/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Property reference for xcsh_site_upgrade_status.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
EnumExtractionComplete: true
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="schema-expected_software_version"></a>

### expected_software_version property

Type: `"string"`. Optional.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
EnumExtractionComplete: false
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
EnumExtractionComplete: true
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
EnumExtractionComplete: false
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
