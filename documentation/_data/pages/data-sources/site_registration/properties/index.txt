---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_site_registration."
xcsh_docs: {"aliases": ["site registration"], "body_bytes": 4866, "body_sha256": "sha256:44a7cf7e617d59f10869c9400a86055c23e1339e97777eabfa0635f3c80d7e7b", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registration:reference", "parent_id": "xcsh-docs:data-sources:site_registration:fundamentals", "path": "documentation/data-sources/site_registration/properties/index.md", "product": "distributed-cloud", "provider_name": "site_registration", "provider_schema_digest": "sha256:057968f86e4ef0ae0087dd4d6097131e285b60998d97655ee1a02c00315f6b0f", "provider_type": "data-sources", "registry_anchor": "canonical-2320211032313111-3313013020320031-3112023331221310-0200221133230022-3003012012313303-0233000303323100-2311002300202220-2110132031202120", "registry_path": "docs/guides/data-sources--site_registration--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["cluster name"], "anchor": "schema-cluster_name", "description": "Cluster name the CE registered with, as reported in its passport. Equals `site_name` for a correctly configured site.", "document_id": "xcsh-docs:data-sources:site_registration:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cluster_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["cluster size"], "anchor": "schema-cluster_size", "description": "Number of nodes the CE reported for its cluster (1 for a single-node site, 3 for a three-node site).", "document_id": "xcsh-docs:data-sources:site_registration:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cluster_size"], "syntax": "attribute", "type": "number"}, {"aliases": ["found"], "anchor": "schema-found", "description": "Whether a registration was resolved. `false` (with no error) while the CE has not registered yet — gate an approval's `count` on this.", "document_id": "xcsh-docs:data-sources:site_registration:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["found"], "syntax": "attribute", "type": "bool"}, {"aliases": ["hostname"], "anchor": "schema-hostname", "description": "Node hostname used to pick one registration when a multi-node site has several. Optional for a single-node site; when omitted, the resolved node's hostname is returned here. Hostnames are only unique within a site.", "document_id": "xcsh-docs:data-sources:site_registration:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["hostname"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "Identifier of this lookup: the registration name when one is found, otherwise null.", "document_id": "xcsh-docs:data-sources:site_registration:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["instance id"], "anchor": "schema-instance_id", "description": "Infrastructure instance identifier reported by the CE registration (`get_spec.infra.instance_id`). This distinguishes rebuilt nodes that reuse the same site and hostname.", "document_id": "xcsh-docs:data-sources:site_registration:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["instance_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["name"], "anchor": "schema-name", "description": "Registration name (`r-<uuid>`) to pass to `xcsh_registration_approval`. Null when `found` is `false`.", "document_id": "xcsh-docs:data-sources:site_registration:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["name"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace holding the registrations. Defaults to `system`, where site registrations live.", "document_id": "xcsh-docs:data-sources:site_registration:reference", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["provider type"], "anchor": "schema-provider_type", "description": "Infrastructure provider the CE reported, e.g. `AZURE`, `AWS`, `GCP`, `VMWARE`.", "document_id": "xcsh-docs:data-sources:site_registration:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["provider_type"], "syntax": "attribute", "type": "string"}, {"aliases": ["site name"], "anchor": "schema-site_name", "description": "Name of the F5 XC site whose CE registration should be resolved. Matched against each registration's `get_spec.passport.cluster_name`.", "document_id": "xcsh-docs:data-sources:site_registration:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["state"], "anchor": "schema-state", "description": "Current registration state, e.g. `PENDING` (awaiting approval) or `ONLINE` (node admitted and healthy).", "document_id": "xcsh-docs:data-sources:site_registration:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["state"], "syntax": "attribute", "type": "string"}, {"aliases": ["uid"], "anchor": "schema-uid", "description": "Unique identifier of the registration (the `<uuid>` part of the name).", "document_id": "xcsh-docs:data-sources:site_registration:reference", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["uid"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registration/properties/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Property reference for xcsh_site_registration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_site_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registration/)
- Property reference

## Direct properties

<a id="schema-cluster_name"></a>

### cluster_name property

Type: `"string"`. Computed.

Cluster name the CE registered with, as reported in its passport. Equals \`site\_name\` for a
correctly configured site.

<a id="schema-cluster_size"></a>

### cluster_size property

Type: `"number"`. Computed.

Number of nodes the CE reported for its cluster (1 for a single-node site, 3 for a three-node site).

<a id="schema-found"></a>

### found property

Type: `"bool"`. Computed.

Whether a registration was resolved. \`false\` (with no error) while the CE has not registered yet —
gate an approval's \`count\` on this.

<a id="schema-hostname"></a>

### hostname property

Type: `"string"`. Optional, Computed.

Node hostname used to pick one registration when a multi-node site has several. Optional for a
single-node site; when omitted, the resolved node's hostname is returned here. Hostnames are only
unique within a site.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

Identifier of this lookup: the registration name when one is found, otherwise null.

<a id="schema-instance_id"></a>

### instance_id property

Type: `"string"`. Computed.

Infrastructure instance identifier reported by the CE registration
(\`get\_spec.infra.instance\_id\`). This distinguishes rebuilt nodes that reuse the same site and
hostname.

<a id="schema-name"></a>

### name property

Type: `"string"`. Computed.

Registration name (\`r-&lt;uuid&gt;\`) to pass to \`xcsh\_registration\_approval\`. Null when
\`found\` is \`false\`.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Optional, Computed.

Namespace holding the registrations. Defaults to \`system\`, where site registrations live.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthAtLeast(1),
}
```

<a id="schema-provider_type"></a>

### provider_type property

Type: `"string"`. Computed.

Infrastructure provider the CE reported, e.g. \`AZURE\`, \`AWS\`, \`GCP\`, \`VMWARE\`.

<a id="schema-site_name"></a>

### site_name property

Type: `"string"`. Required.

Name of the F5 XC site whose CE registration should be resolved. Matched against each registration's
\`get\_spec.passport.cluster\_name\`.

<a id="schema-state"></a>

### state property

Type: `"string"`. Computed.

Current registration state, e.g. \`PENDING\` (awaiting approval) or \`ONLINE\` (node admitted and
healthy).

<a id="schema-uid"></a>

### uid property

Type: `"string"`. Computed.

Unique identifier of the registration (the \`&lt;uuid&gt;\` part of the name).

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `cluster_name` | [cluster_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registration/properties/#schema-cluster_name) |
| `cluster_size` | [cluster_size](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registration/properties/#schema-cluster_size) |
| `found` | [found](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registration/properties/#schema-found) |
| `hostname` | [hostname](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registration/properties/#schema-hostname) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registration/properties/#schema-id) |
| `instance_id` | [instance_id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registration/properties/#schema-instance_id) |
| `name` | [name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registration/properties/#schema-name) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registration/properties/#schema-namespace) |
| `provider_type` | [provider_type](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registration/properties/#schema-provider_type) |
| `site_name` | [site_name](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registration/properties/#schema-site_name) |
| `state` | [state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registration/properties/#schema-state) |
| `uid` | [uid](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registration/properties/#schema-uid) |
