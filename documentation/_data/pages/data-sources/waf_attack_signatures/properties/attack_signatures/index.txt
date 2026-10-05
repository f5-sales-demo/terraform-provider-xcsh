---
page_title: "attack_signatures"
subcategory: ""
description: "List of all supported attack signatures."
xcsh_docs: {"aliases": ["attack signatures"], "body_bytes": 2405, "body_sha256": "sha256:d2278451cbb72b88ca5bb36b6b2b61c3b9ddb790573532489a0b219753911863", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_attack_signatures:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_attack_signatures:properties:attack_signatures", "parent_id": "xcsh-docs:data-sources:waf_attack_signatures:reference", "path": "documentation/data-sources/waf_attack_signatures/properties/attack_signatures/index.md", "product": "distributed-cloud", "provider_name": "waf_attack_signatures", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3013010003301122-1212130103010300-2022332230330221-0011111122102211-0221111032211013-2333310010033331-2210310222101012-3310131002201212", "registry_path": "docs/guides/data-sources--waf_attack_signatures--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["attack_signatures"], "schema_version": 1, "sections": [{"aliases": ["attack signatures accuracy"], "anchor": "schema-attack_signatures--accuracy", "description": "Accuracy. The Signature Accuracy.", "document_id": "xcsh-docs:data-sources:waf_attack_signatures:properties:attack_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["attack_signatures", "accuracy"], "syntax": "attribute", "type": "string"}, {"aliases": ["attack signatures applies to"], "anchor": "schema-attack_signatures--applies_to", "description": "Applies To. The Signature Applies to.", "document_id": "xcsh-docs:data-sources:waf_attack_signatures:properties:attack_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["attack_signatures", "applies_to"], "syntax": "attribute", "type": "string"}, {"aliases": ["attack signatures attack type"], "anchor": "schema-attack_signatures--attack_type", "description": "Attack Type. The Signature Attack Type.", "document_id": "xcsh-docs:data-sources:waf_attack_signatures:properties:attack_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["attack_signatures", "attack_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["attack signatures description spec"], "anchor": "schema-attack_signatures--description_spec", "description": "Description. The Signature Description.", "document_id": "xcsh-docs:data-sources:waf_attack_signatures:properties:attack_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["attack_signatures", "description_spec"], "syntax": "attribute", "type": "string"}, {"aliases": ["attack signatures id"], "anchor": "schema-attack_signatures--id", "description": "ID. The Signature ID.", "document_id": "xcsh-docs:data-sources:waf_attack_signatures:properties:attack_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["attack_signatures", "id"], "syntax": "attribute", "type": "string"}, {"aliases": ["attack signatures last update"], "anchor": "schema-attack_signatures--last_update", "description": "Last Update. The Signature last update time.", "document_id": "xcsh-docs:data-sources:waf_attack_signatures:properties:attack_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["attack_signatures", "last_update"], "syntax": "attribute", "type": "string"}, {"aliases": ["attack signatures name"], "anchor": "schema-attack_signatures--name", "description": "Name. The Signature Name.", "document_id": "xcsh-docs:data-sources:waf_attack_signatures:properties:attack_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["attack_signatures", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["attack signatures references"], "anchor": "schema-attack_signatures--references", "description": "References. The Signature References.", "document_id": "xcsh-docs:data-sources:waf_attack_signatures:properties:attack_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["attack_signatures", "references"], "syntax": "attribute", "type": "list"}, {"aliases": ["attack signatures risk"], "anchor": "schema-attack_signatures--risk", "description": "Risk. The Signature Risk.", "document_id": "xcsh-docs:data-sources:waf_attack_signatures:properties:attack_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["attack_signatures", "risk"], "syntax": "attribute", "type": "string"}, {"aliases": ["attack signatures systems"], "anchor": "schema-attack_signatures--systems", "description": "Systems. The Signature Systems.", "document_id": "xcsh-docs:data-sources:waf_attack_signatures:properties:attack_signatures", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["attack_signatures", "systems"], "syntax": "attribute", "type": "list"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_attack_signatures/properties/attack_signatures/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "List of all supported attack signatures.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# attack_signatures

Breadcrumbs:

- [xcsh_waf_attack_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/properties/)
- attack_signatures

<a id="section"></a>

Type: `"list"`. Computed.

List of all supported attack signatures.

## Direct properties

<a id="schema-attack_signatures--accuracy"></a>

### accuracy property

Type: `"string"`. Computed.

Accuracy. The Signature Accuracy.

<a id="schema-attack_signatures--applies_to"></a>

### applies_to property

Type: `"string"`. Computed.

Applies To. The Signature Applies to.

<a id="schema-attack_signatures--attack_type"></a>

### attack_type property

Type: `"string"`. Computed.

Attack Type. The Signature Attack Type.

<a id="schema-attack_signatures--description_spec"></a>

### description_spec property

Type: `"string"`. Computed.

Description. The Signature Description.

<a id="schema-attack_signatures--id"></a>

### id property

Type: `"string"`. Computed.

ID. The Signature ID.

<a id="schema-attack_signatures--last_update"></a>

### last_update property

Type: `"string"`. Computed.

Last Update. The Signature last update time.

<a id="schema-attack_signatures--name"></a>

### name property

Type: `"string"`. Computed.

Name. The Signature Name.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

<a id="schema-attack_signatures--references"></a>

### references property

Type: `["list", "string"]`. Computed.

References. The Signature References.

<a id="schema-attack_signatures--risk"></a>

### risk property

Type: `"string"`. Computed.

Risk. The Signature Risk.

<a id="schema-attack_signatures--systems"></a>

### systems property

Type: `["list", "string"]`. Computed.

Systems. The Signature Systems.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/properties/)
- [xcsh_waf_attack_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/)
