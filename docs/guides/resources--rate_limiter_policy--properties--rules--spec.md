---
page_title: "rules.spec"
subcategory: "Security"
description: "rules.spec for xcsh_rate_limiter_policy."
xcsh_docs: {"aliases": [], "body_bytes": 5783, "body_sha256": "sha256:6d0e74becdf36a80aa14e0b8bf41c401111a5f1e38be1ff27b2567ac605ba014", "canonical_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec", "child_ids": ["xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:any_asn", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:any_country", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:any_ip", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:apply_rate_limiter", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:asn_list", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:asn_matcher", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:bypass_rate_limiter", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:country_list", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:custom_rate_limiter", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:domain_matcher", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:headers", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:http_method", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:ip_matcher", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:ip_prefix_list", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:path", "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec:segment_policy"], "collection_id": "xcsh-docs:resources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter_policy:properties:rules:spec", "parent_id": "xcsh-docs:resources:rate_limiter_policy:properties:rules", "path": "docs/guides/resources--rate_limiter_policy--properties--rules--spec.md", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["rules", "spec"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter_policy/properties/rules/spec/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "rules.spec for xcsh_rate_limiter_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rules.spec

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md)
- [Property reference](resources--rate_limiter_policy--reference.md)
- [rules](resources--rate_limiter_policy--properties--rules.md)
- rules.spec

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Rate Limiter Rule Specification. Shape of Rate Limiter Rule.

Upstream description:

Shape of Rate Limiter Rule.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_asn",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_asn",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_country",
    "country_list"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("apply_rate_limiter",
    "bypass_rate_limiter"),
  validators.ConflictingObjectAttributes("apply_rate_limiter",
    "custom_rate_limiter"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("bypass_rate_limiter",
    "custom_rate_limiter"),
  validators.ConflictingObjectAttributes("ip_matcher",
    "ip_prefix_list")}
```

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

Terraform syntax:

```terraform
spec {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any_asn](resources--rate_limiter_policy--properties--rules--spec--any_asn.md): complete subsection reference.

- [any_country](resources--rate_limiter_policy--properties--rules--spec--any_country.md): complete subsection reference.

- [any_ip](resources--rate_limiter_policy--properties--rules--spec--any_ip.md): complete subsection reference.

- [apply_rate_limiter](resources--rate_limiter_policy--properties--rules--spec--apply_rate_limiter.md): complete subsection reference.

- [asn_list](resources--rate_limiter_policy--properties--rules--spec--asn_list.md): complete subsection reference.

- [asn_matcher](resources--rate_limiter_policy--properties--rules--spec--asn_matcher.md): complete subsection reference.

- [bypass_rate_limiter](resources--rate_limiter_policy--properties--rules--spec--bypass_rate_limiter.md): complete subsection reference.

- [country_list](resources--rate_limiter_policy--properties--rules--spec--country_list.md): complete subsection reference.

- [custom_rate_limiter](resources--rate_limiter_policy--properties--rules--spec--custom_rate_limiter.md): complete subsection reference.

- [domain_matcher](resources--rate_limiter_policy--properties--rules--spec--domain_matcher.md): complete subsection reference.

- [headers](resources--rate_limiter_policy--properties--rules--spec--headers.md): complete subsection reference.

- [http_method](resources--rate_limiter_policy--properties--rules--spec--http_method.md): complete subsection reference.

- [ip_matcher](resources--rate_limiter_policy--properties--rules--spec--ip_matcher.md): complete subsection reference.

- [ip_prefix_list](resources--rate_limiter_policy--properties--rules--spec--ip_prefix_list.md): complete subsection reference.

- [path](resources--rate_limiter_policy--properties--rules--spec--path.md): complete subsection reference.

- [segment_policy](resources--rate_limiter_policy--properties--rules--spec--segment_policy.md): complete subsection reference.

## Next pages

- [rules.spec.any_asn](resources--rate_limiter_policy--properties--rules--spec--any_asn.md)
- [rules.spec.any_country](resources--rate_limiter_policy--properties--rules--spec--any_country.md)
- [rules.spec.any_ip](resources--rate_limiter_policy--properties--rules--spec--any_ip.md)
- [rules.spec.apply_rate_limiter](resources--rate_limiter_policy--properties--rules--spec--apply_rate_limiter.md)
- [rules.spec.asn_list](resources--rate_limiter_policy--properties--rules--spec--asn_list.md)
- [rules.spec.asn_matcher](resources--rate_limiter_policy--properties--rules--spec--asn_matcher.md)
- [rules.spec.bypass_rate_limiter](resources--rate_limiter_policy--properties--rules--spec--bypass_rate_limiter.md)
- [rules.spec.country_list](resources--rate_limiter_policy--properties--rules--spec--country_list.md)
- [rules.spec.custom_rate_limiter](resources--rate_limiter_policy--properties--rules--spec--custom_rate_limiter.md)
- [rules.spec.domain_matcher](resources--rate_limiter_policy--properties--rules--spec--domain_matcher.md)
- [rules.spec.headers](resources--rate_limiter_policy--properties--rules--spec--headers.md)
- [rules.spec.http_method](resources--rate_limiter_policy--properties--rules--spec--http_method.md)
- [rules.spec.ip_matcher](resources--rate_limiter_policy--properties--rules--spec--ip_matcher.md)
- [rules.spec.ip_prefix_list](resources--rate_limiter_policy--properties--rules--spec--ip_prefix_list.md)
- [rules.spec.path](resources--rate_limiter_policy--properties--rules--spec--path.md)
- [rules.spec.segment_policy](resources--rate_limiter_policy--properties--rules--spec--segment_policy.md)
- [rules](resources--rate_limiter_policy--properties--rules.md)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md)
