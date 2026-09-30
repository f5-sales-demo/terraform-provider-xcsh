---
page_title: "ns_list"
subcategory: ""
description: "ns_list for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 1987, "body_sha256": "sha256:45bc4a48370083c83d620687828a2387c044e2b20f503748f3cf06b555c4dc5f", "canonical_id": "xcsh-docs:resources:global_log_receiver:properties:ns_list", "child_ids": [], "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:ns_list", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "docs/guides/resources--global_log_receiver--properties--ns_list.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ns_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/ns_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ns_list for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ns_list

Breadcrumbs:

- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
- [Property reference](resources--global_log_receiver--reference.md)
- ns_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Namespace List. Namespace List.

Upstream description:

Namespace List.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("namespaces")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
ns_list {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ns_list--namespaces"></a>

### namespaces property

Type: `["list", "string"]`. Optional.

Namespaces. List of namespaces to stream logs for.

Upstream description:

List of namespaces to stream logs for.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeAtMost(16),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 16,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 16,
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
    "ves.io.schema.rules.repeated.max_items": "16"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "16"
  }
}
```

## Next pages

- [Property reference](resources--global_log_receiver--reference.md)
- [xcsh_global_log_receiver](../resources/global_log_receiver.md)
