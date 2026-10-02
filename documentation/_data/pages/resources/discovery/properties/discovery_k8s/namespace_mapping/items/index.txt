---
page_title: "discovery_k8s.namespace_mapping.items"
subcategory: ""
description: "Map K8s namespace(s) to App Namespaces. In Shared Configuration, Discovered Services can only be mapped to a single App Namespace, which is determined by the first matched regex."
xcsh_docs: {"aliases": ["discovery k8s namespace mapping items"], "body_bytes": 4752, "body_sha256": "sha256:07927b7f74fa5fcb2256d165cf9a862a3d2cb6a4ea44862fb036088ffc465d24", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_k8s:namespace_mapping:items", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:namespace_mapping", "path": "documentation/resources/discovery/properties/discovery_k8s/namespace_mapping/items/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3002302011033021-2032112031301321-3113010021213321-1211130323222133-0222322210032313-2033311001220112-2322102303032012-1211323103133123", "registry_path": "docs/guides/resources--discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_k8s", "namespace_mapping", "items"], "schema_version": 1, "sections": [{"aliases": ["namespace"], "anchor": "schema-discovery_k8s--namespace_mapping--items--namespace", "description": "Select a namespace.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:namespace_mapping:items", "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "namespace_mapping", "items", "namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace regex"], "anchor": "schema-discovery_k8s--namespace_mapping--items--namespace_regex", "description": "The regex here will be used to match K8s namespace(s).", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:namespace_mapping:items", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "namespace_mapping", "items", "namespace_regex"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_k8s/namespace_mapping/items/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Map K8s namespace(s) to App Namespaces. In Shared Configuration, Discovered Services can only be mapped to a single App Namespace, which is determined by the first matched regex.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s.namespace_mapping.items

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/)
- [discovery_k8s.namespace_mapping](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/namespace_mapping/)
- discovery_k8s.namespace_mapping.items

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Map K8s namespace(s) to App Namespaces. In Shared Configuration, Discovered Services can only be
mapped to a single App Namespace, which is determined by the first matched regex.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
items {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-discovery_k8s--namespace_mapping--items--namespace"></a>

### namespace property

Type: `"string"`. Optional, Computed.

F5XC Application Namespaces. Select a namespace.

Upstream description:

Select a namespace.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 63),
  stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([-a-z0-9]*[a-z0-9])?$`),
    ""),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-discovery_k8s--namespace_mapping--items--namespace_regex"></a>

### namespace_regex property

Type: `"string"`. Optional.

The regex here will be used to match K8s namespace(s).

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(256),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.regex": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.regex": "true"
  }
}
```

## Next pages

- [discovery_k8s.namespace_mapping](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/namespace_mapping/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
