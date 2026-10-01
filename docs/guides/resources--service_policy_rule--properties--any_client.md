---
page_title: "any_client"
subcategory: ""
description: "any_client for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1582, "body_sha256": "sha256:0874a4d6280867dffbe8f605cb1761a89125f233a3ce6aca9400f668e34e648e", "canonical_id": "xcsh-docs:resources:service_policy_rule:properties:any_client", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:any_client", "parent_id": "xcsh-docs:resources:service_policy_rule:reference", "path": "docs/guides/resources--service_policy_rule--properties--any_client.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["any_client"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/any_client/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "any_client for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# any_client

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
- [Property reference](resources--service_policy_rule--reference.md)
- any_client

<a id="section"></a>

Type: `["object", {}]`. Optional.

\[OneOf: any\_client, client\_name, client\_name\_matcher, client\_selector,
ip\_threat\_category\_list\] Enable this option

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

OneOf alternatives in this subsection:

- [any_client](resources--service_policy_rule--properties--any_client.md#section)
- [client_name](resources--service_policy_rule--reference.md#schema-client_name)
- [client_name_matcher](resources--service_policy_rule--properties--client_name_matcher.md#section)
- [client_selector](resources--service_policy_rule--properties--client_selector.md#section)
- [ip_threat_category_list](resources--service_policy_rule--properties--ip_threat_category_list.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
any_client = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [Property reference](resources--service_policy_rule--reference.md)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
