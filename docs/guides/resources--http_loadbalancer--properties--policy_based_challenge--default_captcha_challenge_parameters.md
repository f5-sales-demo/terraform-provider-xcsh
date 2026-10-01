---
page_title: "policy_based_challenge.default_captcha_challenge_parameters"
subcategory: "Load Balancing"
description: "policy_based_challenge.default_captcha_challenge_parameters for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1191, "body_sha256": "sha256:57d78a0e3bce6db4d3fb7b3e4b1d8675a87e4c90c07dd45446cbd33fce1ba500", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:default_captcha_challenge_parameters", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:default_captcha_challenge_parameters", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge", "path": "docs/guides/resources--http_loadbalancer--properties--policy_based_challenge--default_captcha_challenge_parameters.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["policy_based_challenge", "default_captcha_challenge_parameters"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/policy_based_challenge/default_captcha_challenge_parameters/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_based_challenge.default_captcha_challenge_parameters for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.default_captcha_challenge_parameters

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [policy_based_challenge](resources--http_loadbalancer--properties--policy_based_challenge.md)
- policy_based_challenge.default_captcha_challenge_parameters

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default captcha challenge parameters.

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
default_captcha_challenge_parameters = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [policy_based_challenge](resources--http_loadbalancer--properties--policy_based_challenge.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
