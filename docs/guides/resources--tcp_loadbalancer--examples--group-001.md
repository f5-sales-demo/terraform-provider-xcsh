---
page_title: "xcsh_tcp_loadbalancer examples"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_tcp_loadbalancer examples."
---

# xcsh_tcp_loadbalancer examples

<a id="canonical-0323311213221221-0113310031031222-1330121211223312-1203013310022301-2100320220013120-3302231301121302-0132302333323132-0011101322313000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- Examples

<a id="canonical-3223010303300333-2223101111101312-1020030202300112-0332321102123021-3123312120000211-1300231003012121-3331031213110011-3120330021311021"></a>

### Complete configurations for `xcsh_tcp_loadbalancer`

- [All attributes](resources--tcp_loadbalancer--examples--group-001.md#canonical-1301201233302211-1210110022130302-1231201101130010-2300123023001031-3221310221210311-1311101113122321-1220331031023323-0031331021202001): valid configuration.

- [Resource](resources--tcp_loadbalancer--examples--group-001.md#canonical-3030132121321003-2013130021300310-3112112030301121-3102020110020233-2021003331113332-3103233212023320-3122111323100032-3110120202132001): valid configuration.

- [With annotations](resources--tcp_loadbalancer--examples--group-001.md#canonical-3220100320321221-0012022032320233-2212312002132202-3101120132123322-3330111233032310-0122101011131321-1231103213313112-1111012223123330): valid configuration.

- [With description](resources--tcp_loadbalancer--examples--group-001.md#canonical-2332333133312100-2313121003203331-2133101333332332-1022213331101312-0002032003132003-3030312102032312-0030333122031122-0021031112031122): valid configuration.

- [With healthcheck](resources--tcp_loadbalancer--examples--group-001.md#canonical-1213131201102133-2000302012101102-2030113221011230-0312000232013011-0120022300223021-0301113103321321-3100211002203223-1000133203101100): valid configuration.

- [With labels](resources--tcp_loadbalancer--examples--group-001.md#canonical-1222200102003200-0033011311323320-2102221133112331-2123101222301132-0130320320232323-2331331332003330-2010330231103332-0101300200020222): valid configuration.

- [With listen port](resources--tcp_loadbalancer--examples--group-001.md#canonical-1301223213021230-0103300221332112-2210310131020331-0222120223000330-1210132111201310-0001333321020232-2231032313320311-1033021130000002): valid configuration.

<a id="canonical-1301201233302211-1210110022130302-1231201101130010-2300123023001031-3221310221210311-1311101113122321-1220331031023323-0031331021202001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## All attributes example

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Examples](resources--tcp_loadbalancer--examples--group-001.md#canonical-0323311213221221-0113310031031222-1330121211223312-1203013310022301-2100320220013120-3302231301121302-0132302333323132-0011101322313000)
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

<a id="canonical-3030132121321003-2013130021300310-3112112030301121-3102020110020233-2021003331113332-3103233212023320-3122111323100032-3110120202132001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Examples](resources--tcp_loadbalancer--examples--group-001.md#canonical-0323311213221221-0113310031031222-1330121211223312-1203013310022301-2100320220013120-3302231301121302-0132302333323132-0011101322313000)
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

<a id="canonical-3220100320321221-0012022032320233-2212312002132202-3101120132123322-3330111233032310-0122101011131321-1231103213313112-1111012223123330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With annotations example

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Examples](resources--tcp_loadbalancer--examples--group-001.md#canonical-0323311213221221-0113310031031222-1330121211223312-1203013310022301-2100320220013120-3302231301121302-0132302333323132-0011101322313000)
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

<a id="canonical-2332333133312100-2313121003203331-2133101333332332-1022213331101312-0002032003132003-3030312102032312-0030333122031122-0021031112031122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With description example

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Examples](resources--tcp_loadbalancer--examples--group-001.md#canonical-0323311213221221-0113310031031222-1330121211223312-1203013310022301-2100320220013120-3302231301121302-0132302333323132-0011101322313000)
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

<a id="canonical-1213131201102133-2000302012101102-2030113221011230-0312000232013011-0120022300223021-0301113103321321-3100211002203223-1000133203101100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With healthcheck example

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Examples](resources--tcp_loadbalancer--examples--group-001.md#canonical-0323311213221221-0113310031031222-1330121211223312-1203013310022301-2100320220013120-3302231301121302-0132302333323132-0011101322313000)
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

<a id="canonical-1222200102003200-0033011311323320-2102221133112331-2123101222301132-0130320320232323-2331331332003330-2010330231103332-0101300200020222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With labels example

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Examples](resources--tcp_loadbalancer--examples--group-001.md#canonical-0323311213221221-0113310031031222-1330121211223312-1203013310022301-2100320220013120-3302231301121302-0132302333323132-0011101322313000)
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

<a id="canonical-1301223213021230-0103300221332112-2210310131020331-0222120223000330-1210132111201310-0001333321020232-2231032313320311-1033021130000002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With listen port example

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md#canonical-3323211002020332-1021200301211012-2311222223233022-1102201001122211-1023322232300130-0300323320330112-1021022331100101-3210113012102001)
- [Examples](resources--tcp_loadbalancer--examples--group-001.md#canonical-0323311213221221-0113310031031222-1330121211223312-1203013310022301-2100320220013120-3302231301121302-0132302333323132-0011101322313000)
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
