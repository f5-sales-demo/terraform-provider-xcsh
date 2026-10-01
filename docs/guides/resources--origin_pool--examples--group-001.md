---
page_title: "xcsh_origin_pool examples"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool examples."
---

# xcsh_origin_pool examples

<a id="canonical-00d5a40bf23d1aee66cee08aa2ee9086b6a147ed26743f35e58aaf88fafd260c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6d4ecdf4df7b63722cd9bf761e7786ba6ccf6228921e58f14451f8e3b423b191"></a>

## Examples — Examples / 6d79260cf9d5 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- Examples

<a id="canonical-8a138655261b773009c88e39495abdb00e07ca373e008f886edfc359c08d1d41"></a>

## Complete configurations — Examples / 6d79260cf9d5 / 3

- [Labels update](resources--origin_pool--examples--group-001.md#canonical-59f5bf850c1c63313b3cc9b75b60d6f1a7c256fe7698b65c4f19caa4f8144b13): valid configuration.

- [Multiple origins](resources--origin_pool--examples--group-001.md#canonical-ef2842931dfedaf096cce11db3cb1f2fae5225c7f05f6babd0f171d16eeaaa33): valid configuration.

- [Nested labels](resources--origin_pool--examples--group-001.md#canonical-867b5e0c2acbd6ecd530b74e0ac8fc75f286fa518dba2dfa4477909480ac977e): valid configuration.

- [Port](resources--origin_pool--examples--group-001.md#canonical-8f3de23240446cc731e89fdf9dcc67a5a51733666f1e0efa37ac0edf264c3d13): valid configuration.

- [Public ip](resources--origin_pool--examples--group-001.md#canonical-1c72caad9af39441c2c841acf4cb9a30d36413992b69d2a18a3806320136a566): valid configuration.

- [Resource](resources--origin_pool--examples--group-001.md#canonical-8c3e444ab8d66322017d83bd617dcaf7782c28c900c1f6aed6816334f797cf12): valid configuration.

- [With labels](resources--origin_pool--examples--group-001.md#canonical-b567f2f92979d2e952c82ae9b81b242f08291ac14baa200fba8ce4bcc29da78f): valid configuration.

<a id="canonical-39d19515c5f4fc88936fabb60510dc8cd6eb458907d0e3fcfe4f2665311d135f"></a>

## Next pages — Examples / 6d79260cf9d5 / 4

- [Labels update](resources--origin_pool--examples--group-001.md#canonical-59f5bf850c1c63313b3cc9b75b60d6f1a7c256fe7698b65c4f19caa4f8144b13)
- [Multiple origins](resources--origin_pool--examples--group-001.md#canonical-ef2842931dfedaf096cce11db3cb1f2fae5225c7f05f6babd0f171d16eeaaa33)
- [Nested labels](resources--origin_pool--examples--group-001.md#canonical-867b5e0c2acbd6ecd530b74e0ac8fc75f286fa518dba2dfa4477909480ac977e)
- [Port](resources--origin_pool--examples--group-001.md#canonical-8f3de23240446cc731e89fdf9dcc67a5a51733666f1e0efa37ac0edf264c3d13)
- [Public ip](resources--origin_pool--examples--group-001.md#canonical-1c72caad9af39441c2c841acf4cb9a30d36413992b69d2a18a3806320136a566)
- [Resource](resources--origin_pool--examples--group-001.md#canonical-8c3e444ab8d66322017d83bd617dcaf7782c28c900c1f6aed6816334f797cf12)
- [With labels](resources--origin_pool--examples--group-001.md#canonical-b567f2f92979d2e952c82ae9b81b242f08291ac14baa200fba8ce4bcc29da78f)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-59f5bf850c1c63313b3cc9b75b60d6f1a7c256fe7698b65c4f19caa4f8144b13"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-95eddba81e09d2a1811ef9ebab4b3f47ad85d717fd02af899ab533eed745b634"></a>

## Labels update — Labels update / d153bd89ecd1 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Examples](resources--origin_pool--examples--group-001.md#canonical-00d5a40bf23d1aee66cee08aa2ee9086b6a147ed26743f35e58aaf88fafd260c)
- Labels update

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_origin_pool/labels-update.tf`; digest `sha256:d2a6e4ca39384f01c4d80efec2f06fdf2e82a0293580f75d7c78c1f61e96c022`.

```terraform
# LabelsUpdate — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_origin_pool" "test" {
  name      = "example"
  namespace = "system"

  port = 443

  labels = {
    environment = "example-value"
  }

  origin_servers {
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}
```

<a id="canonical-4a9dcd088787e47bdbfdd7f0d356ffe26b8bdde1a31cb9a2811df958ccb74041"></a>

## Next pages — Labels update / d153bd89ecd1 / 3

- [Examples](resources--origin_pool--examples--group-001.md#canonical-00d5a40bf23d1aee66cee08aa2ee9086b6a147ed26743f35e58aaf88fafd260c)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-ef2842931dfedaf096cce11db3cb1f2fae5225c7f05f6babd0f171d16eeaaa33"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-26b59b1dd425fadd62acae11391c9a7850d8af893f80578382fc106d8e724542"></a>

## Multiple origins — Multiple origins / 277e5678d167 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Examples](resources--origin_pool--examples--group-001.md#canonical-00d5a40bf23d1aee66cee08aa2ee9086b6a147ed26743f35e58aaf88fafd260c)
- Multiple origins

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_origin_pool/multiple-origins.tf`; digest `sha256:bd8a184b9c325d4b7310968c1937c2417c67cf0e0a8113c2bf868b7ccb8905ff`.

```terraform
# MultipleOrigins — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_origin_pool" "test" {
  name      = "example"
  namespace = "system"

  port = 443

  origin_servers {
    public_name {
      dns_name = "backend1.example.com"
    }
  }

  origin_servers {
    public_name {
      dns_name = "backend2.example.com"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}
```

<a id="canonical-ecf8905acb915cd3ec9b9b2cc024d675f2bc7c7ac9ad56fd622e428e87202d38"></a>

## Next pages — Multiple origins / 277e5678d167 / 3

- [Examples](resources--origin_pool--examples--group-001.md#canonical-00d5a40bf23d1aee66cee08aa2ee9086b6a147ed26743f35e58aaf88fafd260c)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-867b5e0c2acbd6ecd530b74e0ac8fc75f286fa518dba2dfa4477909480ac977e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a2ecdb0bd836f32843a5caa6f3608625c9fcba62c575d48b95505f4c732c2b95"></a>

## Nested labels — Nested labels / a785fc67a26c / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Examples](resources--origin_pool--examples--group-001.md#canonical-00d5a40bf23d1aee66cee08aa2ee9086b6a147ed26743f35e58aaf88fafd260c)
- Nested labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_origin_pool/nested-labels.tf`; digest `sha256:2a2c2b7a5d041bafee98f9e4a0464b2e4b5c967ac5f5b30fa2f474ff4f6ef535`.

```terraform
# NestedLabels — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_origin_pool" "test" {
  name      = "example"
  namespace = "system"

  port = 8080

  origin_servers {
    public_ip {
      ip = "192.0.2.1"
    }
    labels = {
      "env" = "test"
      "app" = "demo"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}
```

<a id="canonical-e40a8356cda4127e086fa4404672ac54adba9b5c19e9fc849b0cacdf4eeddc6a"></a>

## Next pages — Nested labels / a785fc67a26c / 3

- [Examples](resources--origin_pool--examples--group-001.md#canonical-00d5a40bf23d1aee66cee08aa2ee9086b6a147ed26743f35e58aaf88fafd260c)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-8f3de23240446cc731e89fdf9dcc67a5a51733666f1e0efa37ac0edf264c3d13"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-52c964010f0328fcdae08336aa2914ef3b3b0aeac1c153d7ad9f963292d634dc"></a>

## Port — Port / 3e23cb6b260e / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Examples](resources--origin_pool--examples--group-001.md#canonical-00d5a40bf23d1aee66cee08aa2ee9086b6a147ed26743f35e58aaf88fafd260c)
- Port

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_origin_pool/port.tf`; digest `sha256:3852ca8c70ab22d421232a89d1c20cf56857ad2339faf797aa538bea24c6a139`.

```terraform
# Port — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_origin_pool" "test" {
  name      = "example"
  namespace = "system"

  port = 443

  origin_servers {
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}
```

<a id="canonical-3074f060985251a8dc8f7dee4cb907496acaece358244a662bb1dbb6a0454e81"></a>

## Next pages — Port / 3e23cb6b260e / 3

- [Examples](resources--origin_pool--examples--group-001.md#canonical-00d5a40bf23d1aee66cee08aa2ee9086b6a147ed26743f35e58aaf88fafd260c)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-1c72caad9af39441c2c841acf4cb9a30d36413992b69d2a18a3806320136a566"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-56800e669d35fa3db54e91114872dbcfc01b81eaab7c624c459c4a0930f5f6f9"></a>

## Public ip — Public ip / 14f4ecb4827b / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Examples](resources--origin_pool--examples--group-001.md#canonical-00d5a40bf23d1aee66cee08aa2ee9086b6a147ed26743f35e58aaf88fafd260c)
- Public ip

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_origin_pool/public-ip.tf`; digest `sha256:1dade891c9c99d46213e707ba46387f7c7b5eed90053ec8616b5c928911c7d54`.

```terraform
# PublicIp — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_origin_pool" "test" {
  name      = "example"
  namespace = "system"

  port = 8080

  origin_servers {
    public_ip {
      ip = "192.0.2.1"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}
```

<a id="canonical-1324b52966cb102cdb3b6774a0c550607f77539965d0f709d71bff10208bbf0f"></a>

## Next pages — Public ip / 14f4ecb4827b / 3

- [Examples](resources--origin_pool--examples--group-001.md#canonical-00d5a40bf23d1aee66cee08aa2ee9086b6a147ed26743f35e58aaf88fafd260c)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-8c3e444ab8d66322017d83bd617dcaf7782c28c900c1f6aed6816334f797cf12"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c44b76a378f68fea439544783c2606d552caf4210b6d8194262447be449c1a4b"></a>

## Resource — Resource / da58e49a005a / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Examples](resources--origin_pool--examples--group-001.md#canonical-00d5a40bf23d1aee66cee08aa2ee9086b6a147ed26743f35e58aaf88fafd260c)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_origin_pool/resource.tf`; digest `sha256:d9f0fe93931beb085130f6173ace667305af3adbe1e9448009e07911a7af8fa0`.

```terraform
# OriginPool Resource Example
# Manages a Origin Pool resource in F5 Distributed Cloud for defining backend server pools for load balancer targets.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic OriginPool configuration
resource "xcsh_origin_pool" "example" {
  name      = "example-origin-pool"
  namespace = "staging"
}
```

<a id="canonical-d60ff5b5d36a2586df9414c8d1827cd6f26398227afb0e7dc85ab3f190be8b88"></a>

## Next pages — Resource / da58e49a005a / 3

- [Examples](resources--origin_pool--examples--group-001.md#canonical-00d5a40bf23d1aee66cee08aa2ee9086b6a147ed26743f35e58aaf88fafd260c)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)

<a id="canonical-b567f2f92979d2e952c82ae9b81b242f08291ac14baa200fba8ce4bcc29da78f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0a600f140fb8823702521b3e746d7c2356f24a36695f9fa1daefe7d848ea2a0b"></a>

## With labels — With labels / cd4c65e701a5 / 2

Breadcrumbs:

- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
- [Examples](resources--origin_pool--examples--group-001.md#canonical-00d5a40bf23d1aee66cee08aa2ee9086b6a147ed26743f35e58aaf88fafd260c)
- With labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_origin_pool/with-labels.tf`; digest `sha256:492fffafa075a1f28ad253873263057c634ae541eb762e94f4e227737c234dcd`.

```terraform
# WithLabels — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_origin_pool" "test" {
  name        = "example"
  namespace   = "system"
  description = "Test origin pool"

  port = 443

  labels = {
    environment = "test"
    team        = "platform"
  }

  origin_servers {
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}
```

<a id="canonical-fff6fbbfa62956b873064c69538ea09a944d9a989da145497bc9052265fd28d5"></a>

## Next pages — With labels / cd4c65e701a5 / 3

- [Examples](resources--origin_pool--examples--group-001.md#canonical-00d5a40bf23d1aee66cee08aa2ee9086b6a147ed26743f35e58aaf88fafd260c)
- [xcsh_origin_pool](../resources/origin_pool.md#canonical-70a9b944294f63f3d3a21174368909ae9ef4e3a2d26f36edb2cb07599127f7aa)
