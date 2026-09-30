---
page_title: "rules.spec.segment_policy.src_any"
subcategory: "Security"
description: "rules.spec.segment_policy.src_any for xcsh_rate_limiter_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1133, "body_sha256": "sha256:310885ca5020122577ebc4b5f2b37299cf537cd3c48233872fa3c86317230fbc", "canonical_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:src_any", "child_ids": [], "collection_id": "xcsh-docs:resources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy:src_any", "parent_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy", "path": "docs/guides/resources--rate_limiter_policy--properties--rules--spec--segment_policy--src_any.md", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "spec", "segment_policy", "src_any"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter_policy/properties/rules/spec/segment_policy/src_any/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.spec.segment_policy.src_any for xcsh_rate_limiter_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# rules.spec.segment_policy.src_any

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md)
- [Property reference](resources--rate_limiter_policy--reference.md)
- [rules](resources--rate_limiter_policy--properties--rules.md)
- [rules.spec](resources--rate_limiter_policy--properties--rules--spec.md)
- [rules.spec.segment_policy](resources--rate_limiter_policy--properties--rules--spec--segment_policy.md)
- rules.spec.segment_policy.src_any

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
src_any = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [rules.spec.segment_policy](resources--rate_limiter_policy--properties--rules--spec--segment_policy.md)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md)
