---
page_title: "policy_based_challenge.rule_list.rules.spec"
subcategory: "Load Balancing"
description: "policy_based_challenge.rule_list.rules.spec for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 10905, "body_sha256": "sha256:f512f2b75b9085a6b2d19f7703143ed3e7115db006b0e2e1c2979d1cfc5616ee", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:any_asn", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:any_client", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:any_ip", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:arg_matchers", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:asn_list", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:asn_matcher", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:body_matcher", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:client_selector", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:cookie_matchers", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:disable_challenge", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:domain_matcher", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:enable_captcha_challenge", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:enable_javascript_challenge", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:headers", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:http_method", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_matcher", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_prefix_list", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:path", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params", "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:tls_fingerprint_matcher"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules", "path": "docs/guides/resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_based_challenge.rule_list.rules.spec for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# policy_based_challenge.rule_list.rules.spec

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [policy_based_challenge](resources--http_loadbalancer--properties--policy_based_challenge.md)
- [policy_based_challenge.rule_list](resources--http_loadbalancer--properties--policy_based_challenge--rule_list.md)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules.md)
- policy_based_challenge.rule_list.rules.spec

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Challenge Rule consists of an unordered list of predicates and an action. The predicates are
evaluated against a set of input fields that are extracted from or derived from an L7 request API. A
request API is considered to match the rule if all predicates in the rule evaluate to true for
that..

Upstream description:

A Challenge Rule consists of an unordered list of predicates and an action. The predicates are
evaluated against a set of input fields that are extracted from or derived from an L7 request API. A
request API is considered to match the rule if all predicates in the rule evaluate to true for that
request. Any predicates that are not specified in a rule are implicitly considered to be true. If a
request API matches a challenge rule, the configured challenge is enforced.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any_asn",
    "asn_list"),
  validators.ConflictingObjectAttributes("any_asn",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("any_client",
    "client_selector"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_matcher"),
  validators.ConflictingObjectAttributes("any_ip",
    "ip_prefix_list"),
  validators.ConflictingObjectAttributes("asn_list",
    "asn_matcher"),
  validators.ConflictingObjectAttributes("disable_challenge",
    "enable_captcha_challenge"),
  validators.ConflictingObjectAttributes("disable_challenge",
    "enable_javascript_challenge"),
  validators.ConflictingObjectAttributes("enable_captcha_challenge",
    "enable_javascript_challenge"),
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
  "x-ves-oneof-field-asn_choice": "[\"any_asn\",\"asn_list\",\"asn_matcher\"]",
  "x-ves-oneof-field-challenge_action": "[\"disable_challenge\",\"enable_captcha_challenge\",\"enable_javascript_challenge\"]",
  "x-ves-oneof-field-client_choice": "[\"any_client\",\"client_selector\"]",
  "x-ves-oneof-field-ip_choice": "[\"any_ip\",\"ip_matcher\",\"ip_prefix_list\"]",
  "x-ves-oneof-field-tls_fingerprint_choice": "[\"tls_fingerprint_matcher\"]"
}
```

Terraform syntax:

```terraform
spec {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any_asn](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--any_asn.md): complete subsection reference.

- [any_client](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--any_client.md): complete subsection reference.

- [any_ip](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--any_ip.md): complete subsection reference.

- [arg_matchers](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--arg_matchers.md): complete subsection reference.

- [asn_list](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--asn_list.md): complete subsection reference.

- [asn_matcher](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--asn_matcher.md): complete subsection reference.

- [body_matcher](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--body_matcher.md): complete subsection reference.

- [client_selector](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--client_selector.md): complete subsection reference.

- [cookie_matchers](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--cookie_matchers.md): complete subsection reference.

- [disable_challenge](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--disable_challenge.md): complete subsection reference.

- [domain_matcher](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--domain_matcher.md): complete subsection reference.

- [enable_captcha_challenge](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--enable_captcha_challenge.md): complete subsection reference.

- [enable_javascript_challenge](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--enable_javascript_challenge.md): complete subsection reference.

<a id="schema-policy_based_challenge--rule_list--rules--spec--expiration_timestamp"></a>

### expiration_timestamp property

Type: `"string"`. Optional.

Specifies expiration\_timestamp the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Upstream description:

The expiration\_timestamp is the RFC 3339 format timestamp at which the containing rule is
considered to be logically expired. The rule continues to exist in the configuration but is not
applied anymore.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "format": "date-time",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [headers](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--headers.md): complete subsection reference.

- [http_method](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--http_method.md): complete subsection reference.

- [ip_matcher](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--ip_matcher.md): complete subsection reference.

- [ip_prefix_list](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--ip_prefix_list.md): complete subsection reference.

- [path](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--path.md): complete subsection reference.

- [query_params](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--query_params.md): complete subsection reference.

- [tls_fingerprint_matcher](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--tls_fingerprint_matcher.md): complete subsection reference.

## Next pages

- [policy_based_challenge.rule_list.rules.spec.any_asn](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--any_asn.md)
- [policy_based_challenge.rule_list.rules.spec.any_client](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--any_client.md)
- [policy_based_challenge.rule_list.rules.spec.any_ip](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--any_ip.md)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--arg_matchers.md)
- [policy_based_challenge.rule_list.rules.spec.asn_list](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--asn_list.md)
- [policy_based_challenge.rule_list.rules.spec.asn_matcher](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--asn_matcher.md)
- [policy_based_challenge.rule_list.rules.spec.body_matcher](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--body_matcher.md)
- [policy_based_challenge.rule_list.rules.spec.client_selector](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--client_selector.md)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--cookie_matchers.md)
- [policy_based_challenge.rule_list.rules.spec.disable_challenge](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--disable_challenge.md)
- [policy_based_challenge.rule_list.rules.spec.domain_matcher](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--domain_matcher.md)
- [policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--enable_captcha_challenge.md)
- [policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--enable_javascript_challenge.md)
- [policy_based_challenge.rule_list.rules.spec.headers](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--headers.md)
- [policy_based_challenge.rule_list.rules.spec.http_method](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--http_method.md)
- [policy_based_challenge.rule_list.rules.spec.ip_matcher](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--ip_matcher.md)
- [policy_based_challenge.rule_list.rules.spec.ip_prefix_list](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--ip_prefix_list.md)
- [policy_based_challenge.rule_list.rules.spec.path](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--path.md)
- [policy_based_challenge.rule_list.rules.spec.query_params](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--query_params.md)
- [policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--tls_fingerprint_matcher.md)
- [policy_based_challenge.rule_list.rules](resources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
