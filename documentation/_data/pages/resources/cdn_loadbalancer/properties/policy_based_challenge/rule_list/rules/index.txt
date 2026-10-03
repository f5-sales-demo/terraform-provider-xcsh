---
page_title: "policy_based_challenge.rule_list.rules"
subcategory: "Load Balancing"
description: "Rules that specify the match conditions and challenge type to be launched. When a challenge type is selected to be always enabled, these rules can be used to disable challenge or launch a different challenge for requests that match the specified conditions."
xcsh_docs: {"aliases": ["policy based challenge rule list rules"], "body_bytes": 3315, "body_sha256": "sha256:fd7e9886936c6c58153611d13668d9debb9a049e13e9ac9f11d97b8f67d87c44", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:metadata", "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list", "path": "documentation/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1223122321232310-0113233200030220-3113200331321213-0303003111130002-0012301211302202-1301121100111032-0033320032020011-0203111200122111", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-013.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["policy_based_challenge", "rule_list", "rules"], "schema_version": 1, "sections": [{"aliases": ["policy based challenge rule list rules metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-policy_based_challenge--rule_list--rules--metadata--name", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:metadata", "type": "requires"}], "schema_path": ["policy_based_challenge", "rule_list", "rules", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec"], "anchor": "section", "description": "A Challenge Rule consists of an unordered list of predicates and an action. The predicates are evaluated against a set of input fields that are extracted from or derived from an L7 request API. A request API is considered to match the rule if all predicates in the rule evaluate to true for that request. Any predicates", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:any_asn,asn_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:any_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:any_asn,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:any_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:any_client,client_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:any_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:any_ip,ip_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:any_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:any_ip,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:any_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:any_asn,asn_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:asn_list,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:any_asn,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:asn_list,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:any_client,client_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:client_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:disable_challenge,enable_captcha_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:disable_challenge", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:disable_challenge,enable_javascript_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:disable_challenge", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:disable_challenge,enable_captcha_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:enable_captcha_challenge", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:enable_captcha_challenge,enable_javascript_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:enable_captcha_challenge", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:disable_challenge,enable_javascript_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:enable_javascript_challenge", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:enable_captcha_challenge,enable_javascript_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:enable_javascript_challenge", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:any_ip,ip_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:ip_matcher,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:any_ip,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:ip_matcher,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_prefix_list", "type": "conflicts"}], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Rules that specify the match conditions and challenge type to be launched. When a challenge type is selected to be always enabled, these rules can be used to disable challenge or launch a different challenge for requests that match the specified conditions.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.rule_list.rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [policy_based_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/policy_based_challenge/)
- [policy_based_challenge.rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/)
- policy_based_challenge.rule_list.rules

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Rules that specify the match conditions and challenge type to be launched. When a challenge type is
selected to be always enabled, these rules can be used to disable challenge or launch a different
challenge for requests that match the specified conditions.

Upstream description:

Rules that specify the match conditions and challenge type to be launched. When a challenge type is
selected to be always enabled, these rules can be used to disable challenge or launch a different
challenge for requests that match the specified conditions.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 64,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": false
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "64",
    "ves.io.schema.rules.repeated.unique_metadata_name": "true"
  }
}
```

Terraform syntax:

```terraform
rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/metadata/): complete subsection reference.

- [spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/): complete subsection reference.

## Next pages

- [policy_based_challenge.rule_list.rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/metadata/)
- [policy_based_challenge.rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/)
- [policy_based_challenge.rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/policy_based_challenge/rule_list/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
