---
page_title: "policy_based_challenge.rule_list.rules"
subcategory: "Load Balancing"
description: "Rules that specify the match conditions and challenge type to be launched. When a challenge type is selected to be always enabled, these rules can be used to disable challenge or launch a different challenge for requests that match the specified conditions."
xcsh_docs: {"aliases": ["policy based challenge rule list rules"], "body_bytes": 3327, "body_sha256": "sha256:77059fa942e24f1b5ec82167c10e9eb96569d50e9557ddad703a29bcfcc53553", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:metadata", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list", "path": "documentation/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1010102303301010-2301220132333203-1322002333000232-0113023301002202-3230123102032122-2330120233123002-1222021211332312-0121121333031030", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-022.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["policy_based_challenge", "rule_list", "rules"], "schema_version": 1, "sections": [{"aliases": ["policy based challenge rule list rules metadata"], "anchor": "section", "description": "MessageMetaType is metadata (common attributes) of a message that only certain messages have. This information is propagated to the metadata of a child object that gets created from the containing message during view processing. The information in this type can be specified by user during create and replace APIs.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:metadata", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-policy_based_challenge--rule_list--rules--metadata--name", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.metadata:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:metadata", "type": "requires"}], "schema_path": ["policy_based_challenge", "rule_list", "rules", "metadata"], "syntax": "block", "type": "object"}, {"aliases": ["policy based challenge rule list rules spec"], "anchor": "section", "description": "A Challenge Rule consists of an unordered list of predicates and an action. The predicates are evaluated against a set of input fields that are extracted from or derived from an L7 request API. A request API is considered to match the rule if all predicates in the rule evaluate to true for that request. Any predicates", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:any_asn,asn_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:any_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:any_asn,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:any_asn", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:any_client,client_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:any_client", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:any_ip,ip_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:any_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:any_ip,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:any_ip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:any_asn,asn_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:asn_list,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:asn_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:any_asn,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:asn_list,asn_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:asn_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:any_client,client_selector", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:client_selector", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:disable_challenge,enable_captcha_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:disable_challenge", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:disable_challenge,enable_javascript_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:disable_challenge", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:disable_challenge,enable_captcha_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:enable_captcha_challenge", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:enable_captcha_challenge,enable_javascript_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:enable_captcha_challenge", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:disable_challenge,enable_javascript_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:enable_javascript_challenge", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:enable_captcha_challenge,enable_javascript_challenge", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:enable_javascript_challenge", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:any_ip,ip_matcher", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:ip_matcher,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_matcher", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:any_ip,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_prefix_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "policy_based_challenge.rule_list.rules.spec:ConflictingObjectAttributes:ip_matcher,ip_prefix_list", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_prefix_list", "type": "conflicts"}], "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Rules that specify the match conditions and challenge type to be launched. When a challenge type is selected to be always enabled, these rules can be used to disable challenge or launch a different challenge for requests that match the specified conditions.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.rule_list.rules

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [policy_based_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/)
- [policy_based_challenge.rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/)
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
      "validatedAt": "2026-10-03T05:10:26+00:00"
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

- [metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/metadata/): complete subsection reference.

- [spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/): complete subsection reference.

## Next pages

- [policy_based_challenge.rule_list.rules.metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/metadata/)
- [policy_based_challenge.rule_list.rules.spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/)
- [policy_based_challenge.rule_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
