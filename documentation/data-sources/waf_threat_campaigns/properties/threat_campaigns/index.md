---
page_title: "threat_campaigns"
subcategory: ""
description: "Threat Campaigns. A list of all supported threat campaigns."
xcsh_docs: {"aliases": ["threat campaigns"], "body_bytes": 2196, "body_sha256": "sha256:e9918257845d87c7d7c827313c7277afc860e5be7947afa8105b214bb9a76ed4", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_threat_campaigns:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "parent_id": "xcsh-docs:data-sources:waf_threat_campaigns:reference", "path": "documentation/data-sources/waf_threat_campaigns/properties/threat_campaigns/index.md", "product": "distributed-cloud", "provider_name": "waf_threat_campaigns", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1112100013220222-0231111002331131-0022021333120010-2333303312103312-1131000232000213-1031301222013030-1133011230131331-2322302201201021", "registry_path": "docs/guides/data-sources--waf_threat_campaigns--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["threat_campaigns"], "schema_version": 1, "sections": [{"aliases": ["threat campaigns attack type"], "anchor": "schema-threat_campaigns--attack_type", "description": "Attack Type. The Threat Campaign Attack Type.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["threat_campaigns", "attack_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["threat campaigns description spec"], "anchor": "schema-threat_campaigns--description_spec", "description": "Description. The Threat Campaign Description.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["threat_campaigns", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["threat campaigns id"], "anchor": "schema-threat_campaigns--id", "description": "ID. The Threat Campaign ID.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["threat_campaigns", "id"], "syntax": "attribute", "type": "string"}, {"aliases": ["threat campaigns intent"], "anchor": "schema-threat_campaigns--intent", "description": "Intent. The Threat Campaign Intent.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["threat_campaigns", "intent"], "syntax": "attribute", "type": "string"}, {"aliases": ["threat campaigns last update"], "anchor": "schema-threat_campaigns--last_update", "description": "Last Update. The Threat Campaign last update time.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["threat_campaigns", "last_update"], "syntax": "attribute", "type": "string"}, {"aliases": ["threat campaigns malwares"], "anchor": "schema-threat_campaigns--malwares", "description": "Malwares. The Threat Campaign Malwares.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["threat_campaigns", "malwares"], "syntax": "attribute", "type": "list"}, {"aliases": ["threat campaigns name"], "anchor": "schema-threat_campaigns--name", "description": "Name. The Threat Campaign Name.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["threat_campaigns", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["threat campaigns references"], "anchor": "schema-threat_campaigns--references", "description": "References. The Threat Campaign References.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["threat_campaigns", "references"], "syntax": "attribute", "type": "list"}, {"aliases": ["threat campaigns risk"], "anchor": "schema-threat_campaigns--risk", "description": "Risk. The Threat Campaign Risk.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["threat_campaigns", "risk"], "syntax": "attribute", "type": "string"}, {"aliases": ["threat campaigns systems"], "anchor": "schema-threat_campaigns--systems", "description": "Systems. The Threat Campaign Systems.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["threat_campaigns", "systems"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_threat_campaigns/properties/threat_campaigns/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Threat Campaigns. A list of all supported threat campaigns.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# threat_campaigns

Breadcrumbs:

- [xcsh_waf_threat_campaigns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/properties/)
- threat_campaigns

<a id="section"></a>

Type: `"list"`. Computed.

Threat Campaigns. A list of all supported threat campaigns.

## Direct properties

<a id="schema-threat_campaigns--attack_type"></a>

### attack_type property

Type: `"string"`. Computed.

Attack Type. The Threat Campaign Attack Type.

<a id="schema-threat_campaigns--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Description. The Threat Campaign Description.

<a id="schema-threat_campaigns--id"></a>

### id property

Type: `"string"`. Computed.

ID. The Threat Campaign ID.

<a id="schema-threat_campaigns--intent"></a>

### intent property

Type: `"string"`. Computed.

Intent. The Threat Campaign Intent.

<a id="schema-threat_campaigns--last_update"></a>

### last_update property

Type: `"string"`. Computed.

Last Update. The Threat Campaign last update time.

<a id="schema-threat_campaigns--malwares"></a>

### malwares property

Type: `["list", "string"]`. Computed.

Malwares. The Threat Campaign Malwares.

<a id="schema-threat_campaigns--name"></a>

### name property

Type: `"string"`. Computed.

Name. The Threat Campaign Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-threat_campaigns--references"></a>

### references property

Type: `["list", "string"]`. Computed.

References. The Threat Campaign References.

<a id="schema-threat_campaigns--risk"></a>

### risk property

Type: `"string"`. Computed.

Risk. The Threat Campaign Risk.

<a id="schema-threat_campaigns--systems"></a>

### systems property

Type: `["list", "string"]`. Computed.

Systems. The Threat Campaign Systems.
