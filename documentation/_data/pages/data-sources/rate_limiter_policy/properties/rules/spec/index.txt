---
page_title: "rules.spec"
subcategory: "Security"
description: "rules.spec for xcsh_rate_limiter_policy."
xcsh_docs: {"aliases": [], "body_bytes": 6723, "body_sha256": "sha256:7b92dab9a1c0ba7f8df5d8580b8061f2a2a9898075805109e0f4d2debac443e2", "child_ids": ["xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:any_asn", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:any_country", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:any_ip", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:apply_rate_limiter", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:asn_list", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:asn_matcher", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:bypass_rate_limiter", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:country_list", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:custom_rate_limiter", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:domain_matcher", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:headers", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:http_method", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:ip_matcher", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:ip_prefix_list", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:path", "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec:segment_policy"], "collection_id": "xcsh-docs:data-sources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules:spec", "parent_id": "xcsh-docs:data-sources:rate_limiter_policy:properties:rules", "path": "documentation/data-sources/rate_limiter_policy/properties/rules/spec/index.md", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["rules", "spec"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/rate_limiter_policy/properties/rules/spec/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.spec for xcsh_rate_limiter_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.spec

Breadcrumbs:

- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/)
- rules.spec

<a id="section"></a>

Type: `"single"`. Computed.

Rate Limiter Rule Specification. Shape of Rate Limiter Rule.

Upstream description:

Shape of Rate Limiter Rule.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"apply_rate_limiter\",\"bypass_rate_limiter\",\"custom_rate_limiter\"]",
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-country_choice": "[\"any_country\",\"country_list\"]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]"
}
```

## Direct properties

- [any_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/any_asn/): complete subsection reference.

- [any_country](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/any_country/): complete subsection reference.

- [any_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/any_ip/): complete subsection reference.

- [apply_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/apply_rate_limiter/): complete subsection reference.

- [asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/asn_list/): complete subsection reference.

- [asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/asn_matcher/): complete subsection reference.

- [bypass_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/bypass_rate_limiter/): complete subsection reference.

- [country_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/country_list/): complete subsection reference.

- [custom_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/custom_rate_limiter/): complete subsection reference.

- [domain_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/domain_matcher/): complete subsection reference.

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/headers/): complete subsection reference.

- [http_method](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/http_method/): complete subsection reference.

- [ip_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/ip_matcher/): complete subsection reference.

- [ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/ip_prefix_list/): complete subsection reference.

- [path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/path/): complete subsection reference.

- [segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/segment_policy/): complete subsection reference.

## Next pages

- [rules.spec.any_asn](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/any_asn/)
- [rules.spec.any_country](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/any_country/)
- [rules.spec.any_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/any_ip/)
- [rules.spec.apply_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/apply_rate_limiter/)
- [rules.spec.asn_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/asn_list/)
- [rules.spec.asn_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/asn_matcher/)
- [rules.spec.bypass_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/bypass_rate_limiter/)
- [rules.spec.country_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/country_list/)
- [rules.spec.custom_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/custom_rate_limiter/)
- [rules.spec.domain_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/domain_matcher/)
- [rules.spec.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/headers/)
- [rules.spec.http_method](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/http_method/)
- [rules.spec.ip_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/ip_matcher/)
- [rules.spec.ip_prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/ip_prefix_list/)
- [rules.spec.path](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/path/)
- [rules.spec.segment_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/spec/segment_policy/)
- [rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/properties/rules/)
- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/rate_limiter_policy/)
