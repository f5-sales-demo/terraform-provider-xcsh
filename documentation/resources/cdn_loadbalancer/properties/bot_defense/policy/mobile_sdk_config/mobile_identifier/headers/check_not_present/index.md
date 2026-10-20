---
page_title: "bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["bot defense policy mobile sdk config mobile identifier headers check not present"], "body_bytes": 1883, "body_sha256": "sha256:5b38a72e140da7904f7fd8472a3a53dd3d552787257a0cdd605968b0eb8bde75", "capabilities": ["cdn", "security.bot-defense"], "category": "cdn", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers:check_not_present", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers", "path": "documentation/resources/cdn_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/mobile_identifier/headers/check_not_present/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2030211311121132-0011030311211213-3220302323232103-3121030121233021-0222321201330021-3320122020201230-3330202211302103-0021233131313011", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-008.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "mobile_sdk_config", "mobile_identifier", "headers", "check_not_present"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/mobile_identifier/headers/check_not_present/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.mobile_sdk_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/mobile_identifier/)
- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/mobile_identifier/headers/)
- bot_defense.policy.mobile_sdk_config.mobile_identifier.headers.check_not_present

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for check not present.

Additional upstream details:

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
check_not_present = {}
```

This is an empty object or choice marker. It has no direct properties.
