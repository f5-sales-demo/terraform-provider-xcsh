---
page_title: "ring_hash"
subcategory: "Load Balancing"
description: "ring_hash for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1238, "body_sha256": "sha256:42ea471766a1c37720a76456767997c09d35edd97ee73f7267d94fd20cf5c956", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:ring_hash:hash_policy"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:ring_hash", "parent_id": "xcsh-docs:resources:http_loadbalancer:reference", "path": "docs/guides/resources--http_loadbalancer--properties--ring_hash.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ring_hash"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/ring_hash/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ring_hash for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ring_hash

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- ring_hash

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Hash Policy List. List of hash policy rules.

Upstream description:

List of hash policy rules.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("hash_policy")}
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
ring_hash {
  # Configure direct properties listed below.
}
```

## Direct properties

- [hash_policy](resources--http_loadbalancer--properties--ring_hash--hash_policy.md): complete subsection reference.

## Next pages

- [ring_hash.hash_policy](resources--http_loadbalancer--properties--ring_hash--hash_policy.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
