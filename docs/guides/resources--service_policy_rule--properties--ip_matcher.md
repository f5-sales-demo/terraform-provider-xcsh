---
page_title: "ip_matcher"
subcategory: ""
description: "ip_matcher for xcsh_service_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1815, "body_sha256": "sha256:6ac58f6ea822841e6d92bc5a73f2f03f6a747728376302e56ef39b9a7078aef5", "canonical_id": "xcsh-docs:resources:service_policy_rule:properties:ip_matcher", "child_ids": ["xcsh-docs:resources:service_policy_rule:properties:ip_matcher:prefix_sets"], "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:service_policy_rule:properties:ip_matcher", "parent_id": "xcsh-docs:resources:service_policy_rule:reference", "path": "docs/guides/resources--service_policy_rule--properties--ip_matcher.md", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ip_matcher"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/properties/ip_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ip_matcher for xcsh_service_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ip_matcher

Breadcrumbs:

- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
- [Property reference](resources--service_policy_rule--reference.md)
- ip_matcher

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_matcher {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-ip_matcher--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [prefix_sets](resources--service_policy_rule--properties--ip_matcher--prefix_sets.md): complete subsection reference.

## Next pages

- [ip_matcher.prefix_sets](resources--service_policy_rule--properties--ip_matcher--prefix_sets.md)
- [Property reference](resources--service_policy_rule--reference.md)
- [xcsh_service_policy_rule](../resources/service_policy_rule.md)
