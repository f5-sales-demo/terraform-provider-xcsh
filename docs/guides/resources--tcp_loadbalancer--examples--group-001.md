---
page_title: "xcsh_tcp_loadbalancer examples"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer examples."
---

# xcsh_tcp_loadbalancer examples

<a id="canonical-3bd67a6917d0d36a7c665af6631f42b190e281d8f2b716721ecbfede0547adc0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-eb133c3fab45547648322c163ee526c9dbd9802570b43199fd367505d8f09d49"></a>

## Examples — Examples / a1c7e7ee12fe / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- Examples

<a id="canonical-a2928530d1bbefed3a3aed151efaff281e8dd44e612cbafb0b0ed4af41481a10"></a>

## Complete configurations — Examples / a1c7e7ee12fe / 3

- [All attributes](resources--tcp_loadbalancer--examples--group-001.md#canonical-7186fca56450a7326d851704b06cb04de9d29935754576b968f4d2fb0df49881): valid configuration.

- [Resource](resources--tcp_loadbalancer--examples--group-001.md#canonical-cc799e4387709c34d658cc59d221422f890fd5fed3be62f8da57b40ed4622781): valid configuration.

- [With annotations](resources--tcp_loadbalancer--examples--group-001.md#canonical-e8438e690628ee2fa6d827a2d161e6fafc56f3b41a4457796d4e7dd6551ab6fc): valid configuration.

- [With description](resources--tcp_loadbalancer--examples--group-001.md#canonical-befdfd90b76438fd9f47ffbe4a9fd47602383783ccd923b60cfda35a0935635a): valid configuration.

- [With healthcheck](resources--tcp_loadbalancer--examples--group-001.md#canonical-6776149f80c864528c5e916c3602e1c5182b0ac9315d3e79d09428eb407e3450): valid configuration.

- [With labels](resources--tcp_loadbalancer--examples--group-001.md#canonical-6a8120e00f175ef892a5f5bd9b46ac5e1ce38bbbbdf7e0fc84f2d4fe11c2022a): valid configuration.

- [With listen port](resources--tcp_loadbalancer--examples--group-001.md#canonical-71ae726c13c29f96a4d1d23d2a62b03c6479587401ff922ead3b7e354f25c002): valid configuration.

<a id="canonical-6c6da79c5c5a40e487495a0c923607cd73314e74e027eeaba4d25cf8e83aab9d"></a>

## Next pages — Examples / a1c7e7ee12fe / 4

- [All attributes](resources--tcp_loadbalancer--examples--group-001.md#canonical-7186fca56450a7326d851704b06cb04de9d29935754576b968f4d2fb0df49881)
- [Resource](resources--tcp_loadbalancer--examples--group-001.md#canonical-cc799e4387709c34d658cc59d221422f890fd5fed3be62f8da57b40ed4622781)
- [With annotations](resources--tcp_loadbalancer--examples--group-001.md#canonical-e8438e690628ee2fa6d827a2d161e6fafc56f3b41a4457796d4e7dd6551ab6fc)
- [With description](resources--tcp_loadbalancer--examples--group-001.md#canonical-befdfd90b76438fd9f47ffbe4a9fd47602383783ccd923b60cfda35a0935635a)
- [With healthcheck](resources--tcp_loadbalancer--examples--group-001.md#canonical-6776149f80c864528c5e916c3602e1c5182b0ac9315d3e79d09428eb407e3450)
- [With labels](resources--tcp_loadbalancer--examples--group-001.md#canonical-6a8120e00f175ef892a5f5bd9b46ac5e1ce38bbbbdf7e0fc84f2d4fe11c2022a)
- [With listen port](resources--tcp_loadbalancer--examples--group-001.md#canonical-71ae726c13c29f96a4d1d23d2a62b03c6479587401ff922ead3b7e354f25c002)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-7186fca56450a7326d851704b06cb04de9d29935754576b968f4d2fb0df49881"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bc3a9ab5c132a220592f2787e3c8ebc8c35240603dc655534c6d6b9790b2e281"></a>

## All attributes — All attributes / 492ab66975b6 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Examples](resources--tcp_loadbalancer--examples--group-001.md#canonical-3bd67a6917d0d36a7c665af6631f42b190e281d8f2b716721ecbfede0547adc0)
- All attributes

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_tcp_loadbalancer/all-attributes.tf`; digest `sha256:c99f272f39faefc7efc9228bcb922ee2344646c08b2b5a103a72ff3f509598a0`.

```terraform
# AllAttributes — Acceptance-test-derived Configuration
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
  name      = "example-pool"
  namespace = "system"
  port      = 443

  origin_servers {
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}

resource "xcsh_tcp_loadbalancer" "test" {
  name        = "example"
  namespace   = "system"
  description = "Acceptance test tcp loadbalancer with all attributes"

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  annotations = {
    purpose = "acceptance-testing"
    owner   = "ci-cd"
  }

  domains     = ["example.example.com"]
  listen_port = 443
  tcp         = {}
  sni         = {}

  origin_pools_weights {
    pool {
      name      = xcsh_origin_pool.test.name
      namespace = "system"
    }
    weight = 1
  }

  advertise_on_public_default_vip = {}
}
```

<a id="canonical-ec44d9b0c44d4c5b01626e2ea94258fc3dfd9160b6d7ff133d6a8daa419e31a3"></a>

## Next pages — All attributes / 492ab66975b6 / 3

- [Examples](resources--tcp_loadbalancer--examples--group-001.md#canonical-3bd67a6917d0d36a7c665af6631f42b190e281d8f2b716721ecbfede0547adc0)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-cc799e4387709c34d658cc59d221422f890fd5fed3be62f8da57b40ed4622781"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a88401b18d6645e9c9cf56c961e5460db4201b5f048bacde9e828bbf72bf08c2"></a>

## Resource — Resource / 48756f6550e5 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Examples](resources--tcp_loadbalancer--examples--group-001.md#canonical-3bd67a6917d0d36a7c665af6631f42b190e281d8f2b716721ecbfede0547adc0)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_tcp_loadbalancer/resource.tf`; digest `sha256:ff5551ee3b1058272666fb12f9288f29debebceead8ddd95379ba00b7d0384b2`.

```terraform
# TCPLoadBalancer Resource Example
# Manages a TCP Load Balancer resource in F5 Distributed Cloud for load balancing TCP traffic across origin pools.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic TCPLoadBalancer configuration
resource "xcsh_tcp_loadbalancer" "example" {
  name      = "example-tcp-loadbalancer"
  namespace = "staging"
}
```

<a id="canonical-2f451371fd2ab0aa11ce55753ac428cd29614a4ffe997a917bd1d0a4b73840f2"></a>

## Next pages — Resource / 48756f6550e5 / 3

- [Examples](resources--tcp_loadbalancer--examples--group-001.md#canonical-3bd67a6917d0d36a7c665af6631f42b190e281d8f2b716721ecbfede0547adc0)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-e8438e690628ee2fa6d827a2d161e6fafc56f3b41a4457796d4e7dd6551ab6fc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-816917a7929d7765df6f13e1fcf0ab2885ac915508ad1d741631d96e0b24fdf9"></a>

## With annotations — With annotations / cf36a76b1233 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Examples](resources--tcp_loadbalancer--examples--group-001.md#canonical-3bd67a6917d0d36a7c665af6631f42b190e281d8f2b716721ecbfede0547adc0)
- With annotations

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_tcp_loadbalancer/with-annotations.tf`; digest `sha256:0dfe6d84947deb60106a2de398c8405b9f312819b62b01977611555b205f5beb`.

```terraform
# WithAnnotations — Acceptance-test-derived Configuration
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
  name      = "example-pool"
  namespace = "system"
  port      = 443

  origin_servers {
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}

resource "xcsh_tcp_loadbalancer" "test" {
  name      = "example"
  namespace = "system"

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  annotations = {
    test_key = "example-value"
  }

  domains     = ["example.example.com"]
  listen_port = 443
  tcp         = {}
  sni         = {}

  origin_pools_weights {
    pool {
      name      = xcsh_origin_pool.test.name
      namespace = "system"
    }
    weight = 1
  }

  advertise_on_public_default_vip = {}
}
```

<a id="canonical-9e1bfb46e287be7e9e890afd948bb54079109551f351a82690b3dcfe5e38d3fd"></a>

## Next pages — With annotations / cf36a76b1233 / 3

- [Examples](resources--tcp_loadbalancer--examples--group-001.md#canonical-3bd67a6917d0d36a7c665af6631f42b190e281d8f2b716721ecbfede0547adc0)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-befdfd90b76438fd9f47ffbe4a9fd47602383783ccd923b60cfda35a0935635a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-4697bd5b4c729c39f96f82f2c041793bbdb786236219829654662c70f566a5b5"></a>

## With description — With description / 4b129a9a76e0 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Examples](resources--tcp_loadbalancer--examples--group-001.md#canonical-3bd67a6917d0d36a7c665af6631f42b190e281d8f2b716721ecbfede0547adc0)
- With description

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_tcp_loadbalancer/with-description.tf`; digest `sha256:5684729ca1aa2364ebb7a66f862baa44ce77ded7c47b4875b858bfb0298f3b7a`.

```terraform
# WithDescription — Acceptance-test-derived Configuration
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
  name      = "example-pool"
  namespace = "system"
  port      = 443

  origin_servers {
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}

resource "xcsh_tcp_loadbalancer" "test" {
  name        = "example"
  namespace   = "system"
  description = "example-value"

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  domains     = ["example.example.com"]
  listen_port = 443
  tcp         = {}
  sni         = {}

  origin_pools_weights {
    pool {
      name      = xcsh_origin_pool.test.name
      namespace = "system"
    }
    weight = 1
  }

  advertise_on_public_default_vip = {}
}
```

<a id="canonical-6a110bf47202b780753d1813af95e3b082571de086525c9086770be42657c247"></a>

## Next pages — With description / 4b129a9a76e0 / 3

- [Examples](resources--tcp_loadbalancer--examples--group-001.md#canonical-3bd67a6917d0d36a7c665af6631f42b190e281d8f2b716721ecbfede0547adc0)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-6776149f80c864528c5e916c3602e1c5182b0ac9315d3e79d09428eb407e3450"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b4dbb6e4a77e0ce54459e533484af3d46b40a32d83bbea9480ce7c7f1f396ce2"></a>

## With healthcheck — With healthcheck / 00046cbc0644 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Examples](resources--tcp_loadbalancer--examples--group-001.md#canonical-3bd67a6917d0d36a7c665af6631f42b190e281d8f2b716721ecbfede0547adc0)
- With healthcheck

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_tcp_loadbalancer/with-healthcheck.tf`; digest `sha256:fdae141ffa6040bc54466938c997fa7b252cc16638c171911504af08dee09ab5`.

```terraform
# WithHealthcheck — Acceptance-test-derived Configuration
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

resource "xcsh_healthcheck" "test" {
  name      = "example-hc"
  namespace = "system"

  healthy_threshold   = 3
  unhealthy_threshold = 1
  timeout             = 3
  interval            = 15

  tcp_health_check {}
}

resource "xcsh_origin_pool" "test" {
  name      = "example-pool"
  namespace = "system"
  port      = 443

  origin_servers {
    public_name {
      dns_name = "example.com"
    }
  }

  healthcheck {
    name      = xcsh_healthcheck.test.name
    namespace = "system"
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}

resource "xcsh_tcp_loadbalancer" "test" {
  name      = "example"
  namespace = "system"

  domains     = ["example.example.com"]
  listen_port = 443
  tcp         = {}
  sni         = {}

  origin_pools_weights {
    pool {
      name      = xcsh_origin_pool.test.name
      namespace = "system"
    }
    weight = 1
  }

  advertise_on_public_default_vip = {}
}
```

<a id="canonical-a5fbd5efcdd21c5e674ea12928a35b23b5d2414600486ff00b4c6056cada8d97"></a>

## Next pages — With healthcheck / 00046cbc0644 / 3

- [Examples](resources--tcp_loadbalancer--examples--group-001.md#canonical-3bd67a6917d0d36a7c665af6631f42b190e281d8f2b716721ecbfede0547adc0)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-6a8120e00f175ef892a5f5bd9b46ac5e1ce38bbbbdf7e0fc84f2d4fe11c2022a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e3dd895baa723bcaa232b3dcbbdd2b05e12620038c8f00d0fb8d23cd67df7a1e"></a>

## With labels — With labels / fba6906de005 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Examples](resources--tcp_loadbalancer--examples--group-001.md#canonical-3bd67a6917d0d36a7c665af6631f42b190e281d8f2b716721ecbfede0547adc0)
- With labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_tcp_loadbalancer/with-labels.tf`; digest `sha256:f2114a8360b2fa5421e12f32a5c7a8f107dc1777845ffa37588092caea04d4c1`.

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
  name      = "example-pool"
  namespace = "system"
  port      = 443

  origin_servers {
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}

resource "xcsh_tcp_loadbalancer" "test" {
  name      = "example"
  namespace = "system"

  labels = {
    environment = "example-value"
    managed_by  = "example-description"
  }

  domains     = ["example.example.com"]
  listen_port = 443
  tcp         = {}
  sni         = {}

  origin_pools_weights {
    pool {
      name      = xcsh_origin_pool.test.name
      namespace = "system"
    }
    weight = 1
  }

  advertise_on_public_default_vip = {}
}
```

<a id="canonical-fe6d3fd35dfdddb8f55db8966fe5d6fba9f0434c86738190e01d8ae6499996de"></a>

## Next pages — With labels / fba6906de005 / 3

- [Examples](resources--tcp_loadbalancer--examples--group-001.md#canonical-3bd67a6917d0d36a7c665af6631f42b190e281d8f2b716721ecbfede0547adc0)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)

<a id="canonical-71ae726c13c29f96a4d1d23d2a62b03c6479587401ff922ead3b7e354f25c002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6b04339dd51d6fdc5a4f9dbb1a7610761316157bd98fb89250ba028eab808db1"></a>

## With listen port — With listen port / 00de83ac06d7 / 2

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
- [Examples](resources--tcp_loadbalancer--examples--group-001.md#canonical-3bd67a6917d0d36a7c665af6631f42b190e281d8f2b716721ecbfede0547adc0)
- With listen port

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_tcp_loadbalancer/with-listen-port.tf`; digest `sha256:418207da4ff4973dbac54d60c3cd0f60db919fc512e95339fa402a4a71300827`.

```terraform
# WithListenPort — Acceptance-test-derived Configuration
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
  name      = "example-pool"
  namespace = "system"
  port      = 443

  origin_servers {
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}

resource "xcsh_tcp_loadbalancer" "test" {
  name      = "example"
  namespace = "system"

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  domains     = ["example.example.com"]
  listen_port = 443

  tcp = {}
  sni = {}

  origin_pools_weights {
    pool {
      name      = xcsh_origin_pool.test.name
      namespace = "system"
    }
    weight = 1
  }

  advertise_on_public_default_vip = {}
}
```

<a id="canonical-45827a3b01deeb50d8046fbc1b5a33b4f2a54adfb29b2a03d0e1f4edcb5f2fe8"></a>

## Next pages — With listen port / 00de83ac06d7 / 3

- [Examples](resources--tcp_loadbalancer--examples--group-001.md#canonical-3bd67a6917d0d36a7c665af6631f42b190e281d8f2b716721ecbfede0547adc0)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-fb94223e49831946b5aabbca528416a54beaec1c30ef8f16492bd411e45c6481)
