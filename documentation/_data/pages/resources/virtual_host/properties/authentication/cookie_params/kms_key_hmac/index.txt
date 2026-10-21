---
page_title: "authentication.cookie_params.kms_key_hmac"
subcategory: ""
description: "Reference to KMS Key Object."
xcsh_docs: {"aliases": ["authentication cookie params kms key hmac"], "body_bytes": 1162, "body_sha256": "sha256:0e0873675ade74bae1ebfea97cd7b92ef07495dab7be7747c964056a317ec204", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params:kms_key_hmac", "parent_id": "xcsh-docs:resources:virtual_host:properties:authentication:cookie_params", "path": "documentation/resources/virtual_host/properties/authentication/cookie_params/kms_key_hmac/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2230033003322103-2333330222332211-1312331010131331-3103000120330311-3302023123120232-3232131132303011-0011333001101110-3203320201201122", "registry_path": "docs/guides/resources--virtual_host--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["authentication", "cookie_params", "kms_key_hmac"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/authentication/cookie_params/kms_key_hmac/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Reference to KMS Key Object.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# authentication.cookie_params.kms_key_hmac

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/)
- [authentication.cookie_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/authentication/cookie_params/)
- authentication.cookie_params.kms_key_hmac

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for kms key hmac.

Additional upstream details:

Reference to KMS Key Object.

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
kms_key_hmac = {}
```

This is an empty object or choice marker. It has no direct properties.
