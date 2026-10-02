---
page_title: "enabled_ssh_access.node_ssh_ports"
subcategory: ""
description: "Enter TCP port and node name per node."
xcsh_docs: {"aliases": ["enabled ssh access node ssh ports"], "body_bytes": 4486, "body_sha256": "sha256:b7538202e30005f32540abe2816ea2e57d5ad707880a9cdcfb5d34e2eb27d6ed", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:node_ssh_ports", "parent_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access", "path": "documentation/resources/nfv_service/properties/enabled_ssh_access/node_ssh_ports/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-2320112001132030-0303302223332210-2023201233202012-2031021302200221-0121132101211303-2310102103021322-3032220021132123-3120320302212102", "registry_path": "docs/guides/resources--nfv_service--reference--group-001.md", "relationships": [{"anchor": "schema-enabled_ssh_access--node_ssh_ports--node_name", "enforcement": "provider-schema", "group": "enabled_ssh_access.node_ssh_ports:RequiredListObjectAttributes:node_name,ssh_port", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:node_ssh_ports", "type": "requires"}, {"anchor": "schema-enabled_ssh_access--node_ssh_ports--ssh_port", "enforcement": "provider-schema", "group": "enabled_ssh_access.node_ssh_ports:RequiredListObjectAttributes:node_name,ssh_port", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:node_ssh_ports", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["enabled_ssh_access", "node_ssh_ports"], "schema_version": 1, "sections": [{"aliases": ["node name"], "anchor": "schema-enabled_ssh_access--node_ssh_ports--node_name", "description": "Node name will be used to match a particular node with the desired TCP port.", "document_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:node_ssh_ports", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enabled_ssh_access", "node_ssh_ports", "node_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["ssh port"], "anchor": "schema-enabled_ssh_access--node_ssh_ports--ssh_port", "description": "Enter TCP port per node.", "document_id": "xcsh-docs:resources:nfv_service:properties:enabled_ssh_access:node_ssh_ports", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["enabled_ssh_access", "node_ssh_ports", "ssh_port"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/enabled_ssh_access/node_ssh_ports/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Enter TCP port and node name per node.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# enabled_ssh_access.node_ssh_ports

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [enabled_ssh_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/enabled_ssh_access/)
- enabled_ssh_access.node_ssh_ports

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Management Node SSH Port. Enter TCP port and node name per node.

Upstream description:

Enter TCP port and node name per node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("node_name",
    "ssh_port")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 2,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 2,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.max_items": "2"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "2"
  }
}
```

Terraform syntax:

```terraform
node_ssh_ports {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-enabled_ssh_access--node_ssh_ports--node_name"></a>

### node_name property

Type: `"string"`. Optional.

Node name will be used to match a particular node with the desired TCP port.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-enabled_ssh_access--node_ssh_ports--ssh_port"></a>

### ssh_port property

Type: `"number"`. Optional.

SSH Port. Enter TCP port per node.

Upstream description:

Enter TCP port per node.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1024, 65535),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 1024
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1024",
    "ves.io.schema.rules.uint32.lte": "65535"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.uint32.gte": "1024",
    "ves.io.schema.rules.uint32.lte": "65535"
  }
}
```

## Next pages

- [enabled_ssh_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/enabled_ssh_access/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
