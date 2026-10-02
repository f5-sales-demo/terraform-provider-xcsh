---
page_title: "threat_campaigns"
subcategory: ""
description: "Threat Campaigns. A list of all supported threat campaigns."
xcsh_docs: {"aliases": ["threat campaigns"], "body_bytes": 2460, "body_sha256": "sha256:22ccb0c1241cee757e5ca9f60e579a8ddc5b2200d81a52fa031f34fa25f59479", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_threat_campaigns:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "parent_id": "xcsh-docs:data-sources:waf_threat_campaigns:reference", "path": "documentation/data-sources/waf_threat_campaigns/properties/threat_campaigns/index.md", "product": "distributed-cloud", "provider_name": "waf_threat_campaigns", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1112100013220222-0231111002331131-0022021333120010-2333303312103312-1131000232000213-1031301222013030-1133011230131331-2322302201201021", "registry_path": "docs/guides/data-sources--waf_threat_campaigns--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["threat_campaigns"], "schema_version": 1, "sections": [{"aliases": ["attack type"], "anchor": "schema-threat_campaigns--attack_type", "description": "Attack Type. The Threat Campaign Attack Type.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["threat_campaigns", "attack_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["description spec"], "anchor": "schema-threat_campaigns--description_spec", "description": "Description. The Threat Campaign Description.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["threat_campaigns", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-threat_campaigns--id", "description": "ID. The Threat Campaign ID.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["threat_campaigns", "id"], "syntax": "attribute", "type": "string"}, {"aliases": ["intent"], "anchor": "schema-threat_campaigns--intent", "description": "Intent. The Threat Campaign Intent.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["threat_campaigns", "intent"], "syntax": "attribute", "type": "string"}, {"aliases": ["last update"], "anchor": "schema-threat_campaigns--last_update", "description": "Last Update. The Threat Campaign last update time.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["threat_campaigns", "last_update"], "syntax": "attribute", "type": "string"}, {"aliases": ["malwares"], "anchor": "schema-threat_campaigns--malwares", "description": "Malwares. The Threat Campaign Malwares.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["threat_campaigns", "malwares"], "syntax": "attribute", "type": "list"}, {"aliases": ["name"], "anchor": "schema-threat_campaigns--name", "description": "Name. The Threat Campaign Name.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["threat_campaigns", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["references"], "anchor": "schema-threat_campaigns--references", "description": "References. The Threat Campaign References.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["threat_campaigns", "references"], "syntax": "attribute", "type": "list"}, {"aliases": ["risk"], "anchor": "schema-threat_campaigns--risk", "description": "Risk. The Threat Campaign Risk.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["threat_campaigns", "risk"], "syntax": "attribute", "type": "string"}, {"aliases": ["systems"], "anchor": "schema-threat_campaigns--systems", "description": "Systems. The Threat Campaign Systems.", "document_id": "xcsh-docs:data-sources:waf_threat_campaigns:properties:threat_campaigns", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["threat_campaigns", "systems"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_threat_campaigns/properties/threat_campaigns/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Threat Campaigns. A list of all supported threat campaigns.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/properties/)
- [xcsh_waf_threat_campaigns](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_threat_campaigns/)
