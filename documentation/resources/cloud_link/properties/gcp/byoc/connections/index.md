---
page_title: "gcp.byoc.connections"
subcategory: ""
description: "Each 'Bring Your Own Connection' represents a virtual connection that the customer has provisioned in the Cloud (example: AWS Direct Connect). F5XC will orchestrate networking resources in the cloud to facilitate seamless private connectivity."
xcsh_docs: {"aliases": ["gcp byoc connections"], "body_bytes": 11324, "body_sha256": "sha256:9f3bd62803e75697bfa616cd6660c0def957c7cefac84cefcbf5b8a5680da691", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections:metadata", "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections:same_as_credential"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections", "parent_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc", "path": "documentation/resources/cloud_link/properties/gcp/byoc/connections/index.md", "product": "distributed-cloud", "provider_name": "cloud_link", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3300300311011020-1232112233202120-3121231302233032-0323330120012132-1203212002211002-0212331312203112-2320223102010211-0101312032120201", "registry_path": "docs/guides/resources--cloud_link--reference--group-001.md", "relationships": [{"anchor": "schema-gcp--byoc--connections--project", "enforcement": "provider-schema", "group": "gcp.byoc.connections:ConflictingListObjectAttributes:project,same_as_credential", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "gcp.byoc.connections:ConflictingListObjectAttributes:project,same_as_credential", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections:same_as_credential", "type": "conflicts"}, {"anchor": "schema-gcp--byoc--connections--interconnect_attachment_name", "enforcement": "provider-schema", "group": "gcp.byoc.connections:RequiredListObjectAttributes:interconnect_attachment_name,region", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections", "type": "requires"}, {"anchor": "schema-gcp--byoc--connections--region", "enforcement": "provider-schema", "group": "gcp.byoc.connections:RequiredListObjectAttributes:interconnect_attachment_name,region", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["gcp", "byoc", "connections"], "schema_version": 1, "sections": [{"aliases": ["gcp byoc connections interconnect attachment name"], "anchor": "schema-gcp--byoc--connections--interconnect_attachment_name", "description": "Name of already-existing GCP Cloud Interconnect Attachment.", "document_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gcp", "byoc", "connections", "interconnect_attachment_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["gcp byoc connections metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-gcp--byoc--connections--metadata--name", "enforcement": "provider-schema", "group": "gcp.byoc.connections.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections:metadata", "type": "requires"}], "schema_path": ["gcp", "byoc", "connections", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["gcp byoc connections project"], "anchor": "schema-gcp--byoc--connections--project", "description": "Exclusive with Specify a GCP Project for the interconnect attachment.", "document_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gcp", "byoc", "connections", "project"], "syntax": "attribute", "type": "string"}, {"aliases": ["gcp byoc connections region"], "anchor": "schema-gcp--byoc--connections--region", "description": "GCP Region in which the GCP Cloud Interconnect attachment is configured.", "document_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gcp", "byoc", "connections", "region"], "syntax": "attribute", "type": "string"}, {"aliases": ["authentication", "credential setup", "credentials", "gcp byoc connections same as credential"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cloud_link:properties:gcp:byoc:connections:same_as_credential", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["gcp", "byoc", "connections", "same_as_credential"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_link/properties/gcp/byoc/connections/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "Each 'Bring Your Own Connection' represents a virtual connection that the customer has provisioned in the Cloud (example: AWS Direct Connect). F5XC will orchestrate networking resources in the cloud to facilitate seamless private connectivity.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# gcp.byoc.connections

Breadcrumbs:

- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/)
- [gcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/)
- [gcp.byoc](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/byoc/)
- gcp.byoc.connections

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Each 'Bring Your Own Connection' represents a virtual connection that the customer has provisioned
in the Cloud (.

Upstream description:

Each 'Bring Your Own Connection' represents a virtual connection that the customer has provisioned
in the Cloud (example: AWS Direct Connect). F5XC will orchestrate networking resources in the cloud
to facilitate seamless private connectivity.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("interconnect_attachment_name",
    "region"),
  validators.ConflictingListObjectAttributes("project",
    "same_as_credential")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 10,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 10,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "10",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

Terraform syntax:

```terraform
connections {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-gcp--byoc--connections--interconnect_attachment_name"></a>

### interconnect_attachment_name property

Type: `"string"`. Optional.

Name of already-existing GCP Cloud Interconnect Attachment.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 63,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "63",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/byoc/connections/metadata/): complete subsection reference.

<a id="schema-gcp--byoc--connections--project"></a>

### project property

Type: `"string"`. Optional.

Exclusive with \[same\_as\_credential\] Specify a GCP Project for the interconnect attachment.

Upstream description:

Exclusive with \[same\_as\_credential\] Specify a GCP Project for the interconnect attachment.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(4, 30),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 30,
  "minLength": 4,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "30",
    "ves.io.schema.rules.string.min_len": "4"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "30",
    "ves.io.schema.rules.string.min_len": "4"
  }
}
```

<a id="schema-gcp--byoc--connections--region"></a>

### region property

Type: `"string"`. Optional.

\[Enum:
asia-east1|asia-east2|asia-northeast1|asia-northeast2|asia-northeast3|asia-southeast1|asia-southeast2|europe-central2|europe-north1|europe-west1|europe-west2|europe-west3|europe-west4|europe-west6|europe-west8|europe-west9|europe-west10|europe-west12|europe-southwest1|me-west1|me-central1|me-central2|northamerica-northeast1|northamerica-northeast2|us-central1|us-east1|us-east4|us-east5|us-south1|us-west1|us-west2|us-west3|us-west4|southamerica-east1|southamerica-west1|australia-southeast1|australia-southeast2|asia-south1|asia-south2\]
GCP Region in which the GCP Cloud Interconnect attachment is configured. Possible values are
\`asia-east1\`, \`asia-east2\`, \`asia-northeast1\`, \`asia-northeast2\`, \`asia-northeast3\`,
\`asia-southeast1\`, \`asia-southeast2\`, \`europe-central2\`, \`europe-north1\`, \`europe-west1\`,
\`europe-west2\`, \`europe-west3\`, \`europe-west4\`, \`europe-west6\`, \`europe-west8\`,
\`europe-west9\`, \`europe-west10\`, \`europe-west12\`, \`europe-southwest1\`, \`me-west1\`,
\`me-central1\`, \`me-central2\`, \`northamerica-northeast1\`, \`northamerica-northeast2\`,
\`us-central1\`, \`us-east1\`, \`us-east4\`, \`us-east5\`, \`us-south1\`, \`us-west1\`,
\`us-west2\`, \`us-west3\`, \`us-west4\`, \`southamerica-east1\`, \`southamerica-west1\`,
\`australia-southeast1\`, \`australia-southeast2\`, \`asia-south1\`, \`asia-south2\`.

Upstream description:

GCP Region in which the GCP Cloud Interconnect attachment is configured.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("asia-east1",
    "asia-east2",
    "asia-northeast1",
    "asia-northeast2",
    "asia-northeast3",
    "asia-southeast1",
    "asia-southeast2",
    "europe-central2",
    "europe-north1",
    "europe-west1",
    "europe-west2",
    "europe-west3",
    "europe-west4",
    "europe-west6",
    "europe-west8",
    "europe-west9",
    "europe-west10",
    "europe-west12",
    "europe-southwest1",
    "me-west1",
    "me-central1",
    "me-central2",
    "northamerica-northeast1",
    "northamerica-northeast2",
    "us-central1",
    "us-east1",
    "us-east4",
    "us-east5",
    "us-south1",
    "us-west1",
    "us-west2",
    "us-west3",
    "us-west4",
    "southamerica-east1",
    "southamerica-west1",
    "australia-southeast1",
    "australia-southeast2",
    "asia-south1",
    "asia-south2"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "asia-east1",
    "asia-east2",
    "asia-northeast1",
    "asia-northeast2",
    "asia-northeast3",
    "asia-southeast1",
    "asia-southeast2",
    "europe-central2",
    "europe-north1",
    "europe-west1",
    "europe-west2",
    "europe-west3",
    "europe-west4",
    "europe-west6",
    "europe-west8",
    "europe-west9",
    "europe-west10",
    "europe-west12",
    "europe-southwest1",
    "me-west1",
    "me-central1",
    "me-central2",
    "northamerica-northeast1",
    "northamerica-northeast2",
    "us-central1",
    "us-east1",
    "us-east4",
    "us-east5",
    "us-south1",
    "us-west1",
    "us-west2",
    "us-west3",
    "us-west4",
    "southamerica-east1",
    "southamerica-west1",
    "australia-southeast1",
    "australia-southeast2",
    "asia-south1",
    "asia-south2"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-03T05:10:26+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"asia-east1\\\",\\\"asia-east2\\\",\\\"asia-northeast1\\\",\\\"asia-northeast2\\\",\\\"asia-northeast3\\\",\\\"asia-southeast1\\\",\\\"asia-southeast2\\\",\\\"europe-central2\\\",\\\"europe-north1\\\",\\\"europe-west1\\\",\\\"europe-west2\\\",\\\"europe-west3\\\",\\\"europe-west4\\\",\\\"europe-west6\\\",\\\"europe-west8\\\",\\\"europe-west9\\\",\\\"europe-west10\\\",\\\"europe-west12\\\",\\\"europe-southwest1\\\",\\\"me-west1\\\",\\\"me-central1\\\",\\\"me-central2\\\",\\\"northamerica-northeast1\\\",\\\"northamerica-northeast2\\\",\\\"us-central1\\\",\\\"us-east1\\\",\\\"us-east4\\\",\\\"us-east5\\\",\\\"us-south1\\\",\\\"us-west1\\\",\\\"us-west2\\\",\\\"us-west3\\\",\\\"us-west4\\\",\\\"southamerica-east1\\\",\\\"southamerica-west1\\\",\\\"australia-southeast1\\\",\\\"australia-southeast2\\\",\\\"asia-south1\\\",\\\"asia-south2\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"asia-east1\\\",\\\"asia-east2\\\",\\\"asia-northeast1\\\",\\\"asia-northeast2\\\",\\\"asia-northeast3\\\",\\\"asia-southeast1\\\",\\\"asia-southeast2\\\",\\\"europe-central2\\\",\\\"europe-north1\\\",\\\"europe-west1\\\",\\\"europe-west2\\\",\\\"europe-west3\\\",\\\"europe-west4\\\",\\\"europe-west6\\\",\\\"europe-west8\\\",\\\"europe-west9\\\",\\\"europe-west10\\\",\\\"europe-west12\\\",\\\"europe-southwest1\\\",\\\"me-west1\\\",\\\"me-central1\\\",\\\"me-central2\\\",\\\"northamerica-northeast1\\\",\\\"northamerica-northeast2\\\",\\\"us-central1\\\",\\\"us-east1\\\",\\\"us-east4\\\",\\\"us-east5\\\",\\\"us-south1\\\",\\\"us-west1\\\",\\\"us-west2\\\",\\\"us-west3\\\",\\\"us-west4\\\",\\\"southamerica-east1\\\",\\\"southamerica-west1\\\",\\\"australia-southeast1\\\",\\\"australia-southeast2\\\",\\\"asia-south1\\\",\\\"asia-south2\\\"]"
  }
}
```

- [same_as_credential](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/byoc/connections/same_as_credential/): complete subsection reference.

## Next pages

- [gcp.byoc.connections.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/byoc/connections/metadata/)
- [gcp.byoc.connections.same_as_credential](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/byoc/connections/same_as_credential/)
- [gcp.byoc](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/properties/gcp/byoc/)
- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_link/)
