---
page_title: "policy_based_challenge.rule_list.rules.spec"
subcategory: "Load Balancing"
description: "policy_based_challenge.rule_list.rules.spec for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 10024, "body_sha256": "sha256:d19406fc6b4410b514b89b460d0e7277dfeba799eb4a80448c4c638e895e4bdc", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:any_asn", "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:any_client", "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:any_ip", "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:arg_matchers", "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:asn_list", "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:asn_matcher", "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:body_matcher", "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:client_selector", "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:cookie_matchers", "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:disable_challenge", "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:domain_matcher", "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:enable_captcha_challenge", "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:enable_javascript_challenge", "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:headers", "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:http_method", "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_matcher", "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:ip_prefix_list", "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:path", "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:query_params", "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec:tls_fingerprint_matcher"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules:spec", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:rule_list:rules", "path": "docs/guides/data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["policy_based_challenge", "rule_list", "rules", "spec"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/policy_based_challenge/rule_list/rules/spec/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "policy_based_challenge.rule_list.rules.spec for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# policy_based_challenge.rule_list.rules.spec

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [policy_based_challenge](data-sources--http_loadbalancer--properties--policy_based_challenge.md)
- [policy_based_challenge.rule_list](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list.md)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules.md)
- policy_based_challenge.rule_list.rules.spec

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [any_asn](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--any_asn.md): complete subsection reference.

- [any_client](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--any_client.md): complete subsection reference.

- [any_ip](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--any_ip.md): complete subsection reference.

- [arg_matchers](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--arg_matchers.md): complete subsection reference.

- [asn_list](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--asn_list.md): complete subsection reference.

- [asn_matcher](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--asn_matcher.md): complete subsection reference.

- [body_matcher](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--body_matcher.md): complete subsection reference.

- [client_selector](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--client_selector.md): complete subsection reference.

- [cookie_matchers](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--cookie_matchers.md): complete subsection reference.

- [disable_challenge](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--disable_challenge.md): complete subsection reference.

- [domain_matcher](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--domain_matcher.md): complete subsection reference.

- [enable_captcha_challenge](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--enable_captcha_challenge.md): complete subsection reference.

- [enable_javascript_challenge](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--enable_javascript_challenge.md): complete subsection reference.

<a id="schema-policy_based_challenge--rule_list--rules--spec--expiration_timestamp"></a>

### expiration_timestamp property

Type: `"string"`. Computed.

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

- [headers](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--headers.md): complete subsection reference.

- [http_method](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--http_method.md): complete subsection reference.

- [ip_matcher](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--ip_matcher.md): complete subsection reference.

- [ip_prefix_list](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--ip_prefix_list.md): complete subsection reference.

- [path](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--path.md): complete subsection reference.

- [query_params](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--query_params.md): complete subsection reference.

- [tls_fingerprint_matcher](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--tls_fingerprint_matcher.md): complete subsection reference.

## Next pages

- [policy_based_challenge.rule_list.rules.spec.any_asn](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--any_asn.md)
- [policy_based_challenge.rule_list.rules.spec.any_client](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--any_client.md)
- [policy_based_challenge.rule_list.rules.spec.any_ip](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--any_ip.md)
- [policy_based_challenge.rule_list.rules.spec.arg_matchers](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--arg_matchers.md)
- [policy_based_challenge.rule_list.rules.spec.asn_list](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--asn_list.md)
- [policy_based_challenge.rule_list.rules.spec.asn_matcher](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--asn_matcher.md)
- [policy_based_challenge.rule_list.rules.spec.body_matcher](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--body_matcher.md)
- [policy_based_challenge.rule_list.rules.spec.client_selector](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--client_selector.md)
- [policy_based_challenge.rule_list.rules.spec.cookie_matchers](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--cookie_matchers.md)
- [policy_based_challenge.rule_list.rules.spec.disable_challenge](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--disable_challenge.md)
- [policy_based_challenge.rule_list.rules.spec.domain_matcher](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--domain_matcher.md)
- [policy_based_challenge.rule_list.rules.spec.enable_captcha_challenge](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--enable_captcha_challenge.md)
- [policy_based_challenge.rule_list.rules.spec.enable_javascript_challenge](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--enable_javascript_challenge.md)
- [policy_based_challenge.rule_list.rules.spec.headers](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--headers.md)
- [policy_based_challenge.rule_list.rules.spec.http_method](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--http_method.md)
- [policy_based_challenge.rule_list.rules.spec.ip_matcher](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--ip_matcher.md)
- [policy_based_challenge.rule_list.rules.spec.ip_prefix_list](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--ip_prefix_list.md)
- [policy_based_challenge.rule_list.rules.spec.path](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--path.md)
- [policy_based_challenge.rule_list.rules.spec.query_params](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--query_params.md)
- [policy_based_challenge.rule_list.rules.spec.tls_fingerprint_matcher](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules--spec--tls_fingerprint_matcher.md)
- [policy_based_challenge.rule_list.rules](data-sources--http_loadbalancer--properties--policy_based_challenge--rule_list--rules.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
