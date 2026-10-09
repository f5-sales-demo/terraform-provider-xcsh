---
page_title: "discovery_k8s.publish_info.publish"
subcategory: ""
description: "K8SPublishType."
xcsh_docs: {"aliases": ["discovery k8s publish info publish"], "body_bytes": 2910, "body_sha256": "sha256:3df5bbe07ccc32a7f313635542a684e6330fd5d056455e59b2cba832e9d9e215", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info", "path": "documentation/resources/discovery/properties/discovery_k8s/publish_info/publish/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-2201132213231022-0230303111210031-0130220302231300-2323030120010232-2210223132002011-2002001223211111-3210202213010231-0112131333302012", "registry_path": "docs/guides/resources--discovery--reference--group-002.md", "relationships": [{"anchor": "schema-discovery_k8s--publish_info--publish--namespace", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info.publish:RequiredObjectAttributes:namespace", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_k8s", "publish_info", "publish"], "schema_version": 1, "sections": [{"aliases": ["discovery k8s publish info publish namespace"], "anchor": "schema-discovery_k8s--publish_info--publish--namespace", "description": "The namespace where the service/endpoints need to be created if it's not included in the domain. The external K8s administrator needs to ensure that the namespace exists.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional", "computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "publish_info", "publish", "namespace"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_k8s/publish_info/publish/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "K8SPublishType.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["discoveryCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s.publish_info.publish

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/)
- [discovery_k8s.publish_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/)
- discovery_k8s.publish_info.publish

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

K8SPublishType.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("namespace")}
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
publish {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-discovery_k8s--publish_info--publish--namespace"></a>

### namespace property

Type: `"string"`. Optional, Computed.

The namespace where the service/endpoints need to be created if it's not included in the domain. The
external K8s administrator needs to ensure that the namespace exists.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 64),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
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
    "maxLength": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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
    "ves.io.schema.rules.string.max_len": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "64"
  }
}
```
