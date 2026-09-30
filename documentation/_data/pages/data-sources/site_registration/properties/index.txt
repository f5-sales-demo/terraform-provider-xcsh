---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_site_registration."
xcsh_docs: {"aliases": [], "body_bytes": 4870, "body_sha256": "sha256:b0df17f43f5f60cdfef490bdceab24b5d49de37139e4d910da6c64c869bb09a2", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registration:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registration:reference", "parent_id": "xcsh-docs:data-sources:site_registration:fundamentals", "path": "documentation/data-sources/site_registration/properties/index.md", "provider_name": "site_registration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registration/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_site_registration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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

## Next pages

- [xcsh_site_registration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registration/)
