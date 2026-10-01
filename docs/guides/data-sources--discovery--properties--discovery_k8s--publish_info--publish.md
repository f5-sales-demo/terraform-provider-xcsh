---
page_title: "discovery_k8s.publish_info.publish"
subcategory: ""
description: "discovery_k8s.publish_info.publish for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 2603, "body_sha256": "sha256:2892165fc15c8aaff2aa99c77b73da27b9c41f5f16afa301d8df8ff0fa65d5f6", "canonical_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:publish", "child_ids": [], "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:publish", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info", "path": "docs/guides/data-sources--discovery--properties--discovery_k8s--publish_info--publish.md", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["discovery_k8s", "publish_info", "publish"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_k8s/publish_info/publish/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_k8s.publish_info.publish for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s.publish_info.publish

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md)
- [Property reference](data-sources--discovery--reference.md)
- [discovery_k8s](data-sources--discovery--properties--discovery_k8s.md)
- [discovery_k8s.publish_info](data-sources--discovery--properties--discovery_k8s--publish_info.md)
- discovery_k8s.publish_info.publish

<a id="section"></a>

Type: `"single"`. Computed.

K8SPublishType.

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

## Direct properties

<a id="schema-discovery_k8s--publish_info--publish--namespace"></a>

### namespace property

Type: `"string"`. Computed.

The namespace where the service/endpoints need to be created if it's not included in the domain. The
external K8s administrator needs to ensure that the namespace exists.

Upstream description:

The namespace where the service/endpoints need to be created if it's not included in the domain. The
external K8s administrator needs to ensure that the namespace exists.

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

## Next pages

- [discovery_k8s.publish_info](data-sources--discovery--properties--discovery_k8s--publish_info.md)
- [xcsh_discovery](../data-sources/discovery.md)
