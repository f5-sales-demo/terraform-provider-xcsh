---
page_title: "xcsh_http_loadbalancer examples"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_http_loadbalancer examples."
---

# xcsh_http_loadbalancer examples

<a id="canonical-2021030002312302-3322220033323000-0123230032002330-3323322031301112-3231001302332112-3232030323010112-1100302103320001-3132301331220203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- Examples

<a id="canonical-3220012202012203-1002212203200111-0030021000231000-0032101303320302-1300203021000102-1002033211110303-2300310132012003-3222323211022202"></a>

### Complete configurations for `xcsh_http_loadbalancer`

- [Conflict protocol](resources--http_loadbalancer--examples--group-001.md#canonical-1223222112033332-2232002203232333-3101301121330131-2213023323032210-0223113110311202-2031310032310212-1003230110132011-0113230321111130): expected conflict.

- [Do not advertise](resources--http_loadbalancer--examples--group-001.md#canonical-3120101203022113-1003321113112233-1120222122222230-3112230201002021-3221132031022310-2103110202300300-0223312230131211-1211010020313032): valid configuration.

- [Https auto cert virtual site](resources--http_loadbalancer--examples--group-001.md#canonical-2102202300103120-3301312001320200-0322123133022133-3201323000313020-2112130222312120-3120103101103003-3000003011113212-2310111232112030): valid configuration.

- [Https auto cert](resources--http_loadbalancer--examples--group-001.md#canonical-3321321033101012-1202003033201201-1331301332133233-3003021113023032-1001111113131231-3000131321113112-1333233213213010-1031301300331301): valid configuration.

- [Ip reputation](resources--http_loadbalancer--examples--group-001.md#canonical-3110122013200000-2033131031130223-0220320122010133-1210113221211021-2022112222220110-3312002203321312-2033013031121321-3030122033013033): valid configuration.

- [Js challenge](resources--http_loadbalancer--examples--group-001.md#canonical-0213133332210331-0002031102312131-1213331311221230-1113232113102322-3210333313011101-3001033001020110-0221323210211211-0031113201312130): valid configuration.

- [Labels update](resources--http_loadbalancer--examples--group-001.md#canonical-1310030100232322-2010311111022213-1201131221102123-0031113301000323-0000022033010223-0233103123001000-1122231311231323-1301311130320302): valid configuration.

- [Least active](resources--http_loadbalancer--examples--group-001.md#canonical-3320011031003002-3333211212322330-0030003110001012-1023320020003123-1221211321102033-3132310113012033-1031032132322023-2031323311321203): valid configuration.

- [Live https auto cert virtual site](resources--http_loadbalancer--examples--group-001.md#canonical-2323223122222222-3210000131013200-0233201102303202-2101213100320203-0110303302122200-0313333001332300-3011033100330122-3110032010230012): valid configuration.

- [Resource](resources--http_loadbalancer--examples--group-001.md#canonical-2130121101331012-0321210331311101-2302201121002303-1022211303332323-0012333220122203-3302020030230303-1033222332021201-2223301113220321): valid configuration.

- [Source ip stickiness](resources--http_loadbalancer--examples--group-001.md#canonical-1123330203031122-0213112100022302-2131233022233300-2211010100313212-3202302233332233-0021210322133220-1232023231200121-3211022211030331): valid configuration.

- [With domains](resources--http_loadbalancer--examples--group-001.md#canonical-0002002311230333-2210103321020131-0331202032212021-2312020110212122-3030132321312222-2312012112230330-2212123212212122-1032110122121202): valid configuration.

- [With labels](resources--http_loadbalancer--examples--group-001.md#canonical-2313131123322031-2332133330031200-0021120210213333-3110220323110302-0113201013010213-2031021310201221-2002232220332111-0100223312320302): valid configuration.

- [With origin pool](resources--http_loadbalancer--examples--group-001.md#canonical-1021031001203102-1232322101103101-3000301221322230-3313031312223200-3223100131230111-0333323031311201-0101001312210011-1013201123131330): valid configuration.

<a id="canonical-1223222112033332-2232002203232333-3101301121330131-2213023323032210-0223113110311202-2031310032310212-1003230110132011-0113230321111130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Conflict protocol

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Examples](resources--http_loadbalancer--examples--group-001.md#canonical-2021030002312302-3322220033323000-0123230032002330-3323322031301112-3231001302332112-3232030323010112-1100302103320001-3132301331220203)
- Conflict protocol

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **expected conflict**.

Source: `examples/resources/xcsh_http_loadbalancer/conflict-protocol.tf`; digest `sha256:536a4332401ed618974908dabf261dd097e7687d8727db003cdf77a23e114526`.

```terraform
# ConflictProtocol — Negative Configuration Example
# Acceptance-test-derived conflict fixture; not a successful configuration.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_http_loadbalancer" "test" {
  name      = "example"
  namespace = "system"
  domains   = ["test.example.com"]

  http {
    port = 80
  }

  https_auto_cert {
    add_hsts                 = false
    no_mtls                  = {}
    default_header           = {}
    enable_path_normalize    = {}
    non_default_loadbalancer = {}
  }

  advertise_on_public_default_vip = {}
}
```

<a id="canonical-3120101203022113-1003321113112233-1120222122222230-3112230201002021-3221132031022310-2103110202300300-0223312230131211-1211010020313032"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Do not advertise example

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Examples](resources--http_loadbalancer--examples--group-001.md#canonical-2021030002312302-3322220033323000-0123230032002330-3323322031301112-3231001302332112-3232030323010112-1100302103320001-3132301331220203)
- Do not advertise

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/do-not-advertise.tf`; digest `sha256:3444dd9792bdb6ac54c1b3867e0fc76032324827010353b443a99ba9f7d0f215`.

```terraform
# DoNotAdvertise — Acceptance-test-derived Configuration
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

resource "xcsh_http_loadbalancer" "test" {
  name      = "example"
  namespace = "system"
  domains   = ["test.example.com"]

  http {
    port = 80
  }

  do_not_advertise = {}
}
```

<a id="canonical-2102202300103120-3301312001320200-0322123133022133-3201323000313020-2112130222312120-3120103101103003-3000003011113212-2310111232112030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Https auto cert virtual site example

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Examples](resources--http_loadbalancer--examples--group-001.md#canonical-2021030002312302-3322220033323000-0123230032002330-3323322031301112-3231001302332112-3232030323010112-1100302103320001-3132301331220203)
- Https auto cert virtual site

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/https-auto-cert-virtual-site.tf`; digest `sha256:88746cf0029b0d4fda1c11f433ad73b5d62d553eb9971ee08c2988ec5e8c543d`.

```terraform
# HttpsAutoCertVirtualSite — Acceptance-test-derived Configuration
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

resource "xcsh_http_loadbalancer" "test" {
  name      = "example"
  namespace = "example-value"
  domains   = ["test.example.com"]

  https_auto_cert {}

  advertise_custom {
    advertise_where {
      virtual_site {
        network = "SITE_NETWORK_INSIDE_AND_OUTSIDE"
        virtual_site {
          name      = "example-description"
          namespace = "example-value"
        }
      }
      use_default_port = {}
    }
  }
}
```

<a id="canonical-3321321033101012-1202003033201201-1331301332133233-3003021113023032-1001111113131231-3000131321113112-1333233213213010-1031301300331301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Https auto cert example

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Examples](resources--http_loadbalancer--examples--group-001.md#canonical-2021030002312302-3322220033323000-0123230032002330-3323322031301112-3231001302332112-3232030323010112-1100302103320001-3132301331220203)
- Https auto cert

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/https-auto-cert.tf`; digest `sha256:0d662f3b7cc004a1f0f8a4fb17d79cd9f87b6670b4fb87c8b2fac5e03df5a58a`.

```terraform
# HttpsAutoCert — Acceptance-test-derived Configuration
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

resource "xcsh_http_loadbalancer" "test" {
  name      = "example"
  namespace = "system"

  domains = ["test.example.com"]

  https_auto_cert {
    add_hsts                 = false
    no_mtls                  = {}
    default_header           = {}
    enable_path_normalize    = {}
    non_default_loadbalancer = {}
  }

  advertise_on_public_default_vip = {}
}
```

<a id="canonical-3110122013200000-2033131031130223-0220320122010133-1210113221211021-2022112222220110-3312002203321312-2033013031121321-3030122033013033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Ip reputation example

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Examples](resources--http_loadbalancer--examples--group-001.md#canonical-2021030002312302-3322220033323000-0123230032002330-3323322031301112-3231001302332112-3232030323010112-1100302103320001-3132301331220203)
- Ip reputation

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/ip-reputation.tf`; digest `sha256:dcf6720963e3d25c8204a3d4b94665f5900cbfd21d3a740b92a547a3fa256c53`.

```terraform
# IpReputation — Acceptance-test-derived Configuration
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

resource "xcsh_http_loadbalancer" "test" {
  name      = "example"
  namespace = "system"
  domains   = ["test.example.com"]

  http {
    port = 80
  }

  enable_ip_reputation {
    ip_threat_categories = ["SPAM_SOURCES"]
  }

  advertise_on_public_default_vip = {}
}
```

<a id="canonical-0213133332210331-0002031102312131-1213331311221230-1113232113102322-3210333313011101-3001033001020110-0221323210211211-0031113201312130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Js challenge example

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Examples](resources--http_loadbalancer--examples--group-001.md#canonical-2021030002312302-3322220033323000-0123230032002330-3323322031301112-3231001302332112-3232030323010112-1100302103320001-3132301331220203)
- Js challenge

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/js-challenge.tf`; digest `sha256:d2d6735f1a5073a1a55a04da2ce4e35a1688038eb16678dc8de5248c6cb894e4`.

```terraform
# JsChallenge — Acceptance-test-derived Configuration
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

resource "xcsh_http_loadbalancer" "test" {
  name      = "example"
  namespace = "system"
  domains   = ["test.example.com"]

  http {
    port = 80
  }

  js_challenge {
    js_script_delay = 5000
    cookie_expiry   = 3600
  }

  advertise_on_public_default_vip = {}
}
```

<a id="canonical-1310030100232322-2010311111022213-1201131221102123-0031113301000323-0000022033010223-0233103123001000-1122231311231323-1301311130320302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Labels update example

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Examples](resources--http_loadbalancer--examples--group-001.md#canonical-2021030002312302-3322220033323000-0123230032002330-3323322031301112-3231001302332112-3232030323010112-1100302103320001-3132301331220203)
- Labels update

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/labels-update.tf`; digest `sha256:047c3e407fe8f84c02b646d4800c1eace6de8efae5137aae3f4ce8865c70e147`.

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

resource "xcsh_http_loadbalancer" "test" {
  name      = "example"
  namespace = "system"

  labels = {
    environment = "example-value"
    managed_by  = "terraform"
  }

  domains = ["test.example.com"]

  http {
    port = 80
  }

  advertise_on_public_default_vip = {}
}
```

<a id="canonical-3320011031003002-3333211212322330-0030003110001012-1023320020003123-1221211321102033-3132310113012033-1031032132322023-2031323311321203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Least active example

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Examples](resources--http_loadbalancer--examples--group-001.md#canonical-2021030002312302-3322220033323000-0123230032002330-3323322031301112-3231001302332112-3232030323010112-1100302103320001-3132301331220203)
- Least active

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/least-active.tf`; digest `sha256:70324220d9c9fa5c1c5615baf4370d340fb061db4d54cc32d3c9ec11039e1ef2`.

```terraform
# LeastActive — Acceptance-test-derived Configuration
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

resource "xcsh_http_loadbalancer" "test" {
  name      = "example"
  namespace = "system"
  domains   = ["test.example.com"]

  http {
    port = 80
  }

  least_active = {}

  advertise_on_public_default_vip = {}
}
```

<a id="canonical-2323223122222222-3210000131013200-0233201102303202-2101213100320203-0110303302122200-0313333001332300-3011033100330122-3110032010230012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Live https auto cert virtual site example

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Examples](resources--http_loadbalancer--examples--group-001.md#canonical-2021030002312302-3322220033323000-0123230032002330-3323322031301112-3231001302332112-3232030323010112-1100302103320001-3132301331220203)
- Live https auto cert virtual site

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/live-https-auto-cert-virtual-site.tf`; digest `sha256:d4a0031e7002f14871ec0e167d4c2e4fbb19bca356c59205380d515763ed3321`.

```terraform
# LiveHTTPSAutoCertVirtualSite — Acceptance-test-derived Configuration
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

resource "xcsh_virtual_site" "test" {
  name      = "example-description"
  namespace = "example-value"
  site_type = "CUSTOMER_EDGE"

  site_selector {
    expressions = ["site_type=customer_edge"]
  }
}

resource "xcsh_http_loadbalancer" "test" {
  depends_on = [xcsh_virtual_site.test]
  name       = "example"
  namespace  = "example-value"
  domains    = ["test.example.com"]

  https_auto_cert {}

  advertise_custom {
    advertise_where {
      virtual_site {
        network = "SITE_NETWORK_INSIDE_AND_OUTSIDE"
        virtual_site {
          name      = xcsh_virtual_site.test.name
          namespace = "example-value"
        }
      }
      use_default_port = {}
    }
  }
}
```

<a id="canonical-2130121101331012-0321210331311101-2302201121002303-1022211303332323-0012333220122203-3302020030230303-1033222332021201-2223301113220321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Examples](resources--http_loadbalancer--examples--group-001.md#canonical-2021030002312302-3322220033323000-0123230032002330-3323322031301112-3231001302332112-3232030323010112-1100302103320001-3132301331220203)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/resource.tf`; digest `sha256:2cc83254a22eded7b491fec9fc0411929e49400732480241f5c8b07d373a95c1`.

```terraform
# HTTPLoadBalancer Resource Example
# Manages a HTTP Load Balancer resource in F5 Distributed Cloud for load balancing HTTP/HTTPS traffic with routing and security controls.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic HTTPLoadBalancer configuration
resource "xcsh_http_loadbalancer" "example" {
  name      = "example-http-loadbalancer"
  namespace = "staging"

  domains = ["example-value"]
}
```

<a id="canonical-1123330203031122-0213112100022302-2131233022233300-2211010100313212-3202302233332233-0021210322133220-1232023231200121-3211022211030331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Source ip stickiness example

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Examples](resources--http_loadbalancer--examples--group-001.md#canonical-2021030002312302-3322220033323000-0123230032002330-3323322031301112-3231001302332112-3232030323010112-1100302103320001-3132301331220203)
- Source ip stickiness

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/source-ip-stickiness.tf`; digest `sha256:739e9d735e5602a3b26f0711cc74f66e68a750409c6a597ce93f240dc10a68a7`.

```terraform
# SourceIpStickiness — Acceptance-test-derived Configuration
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

resource "xcsh_http_loadbalancer" "test" {
  name      = "example"
  namespace = "system"
  domains   = ["test.example.com"]

  http {
    port = 80
  }

  source_ip_stickiness = {}

  advertise_on_public_default_vip = {}
}
```

<a id="canonical-0002002311230333-2210103321020131-0331202032212021-2312020110212122-3030132321312222-2312012112230330-2212123212212122-1032110122121202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With domains example

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Examples](resources--http_loadbalancer--examples--group-001.md#canonical-2021030002312302-3322220033323000-0123230032002330-3323322031301112-3231001302332112-3232030323010112-1100302103320001-3132301331220203)
- With domains

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/with-domains.tf`; digest `sha256:1a4378104ebc071c5ead72e6964e63ea12d8aa818fd0da9d4fb4731e46b0ab68`.

```terraform
# WithDomains — Acceptance-test-derived Configuration
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

resource "xcsh_http_loadbalancer" "test" {
  name      = "example"
  namespace = "system"

  labels = {
    environment = "test"
  }

  domains = [
    "app.example.com",
    "api.example.com"
  ]

  http {
    port = 80
  }

  advertise_on_public_default_vip = {}
}
```

<a id="canonical-2313131123322031-2332133330031200-0021120210213333-3110220323110302-0113201013010213-2031021310201221-2002232220332111-0100223312320302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With labels example

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Examples](resources--http_loadbalancer--examples--group-001.md#canonical-2021030002312302-3322220033323000-0123230032002330-3323322031301112-3231001302332112-3232030323010112-1100302103320001-3132301331220203)
- With labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/with-labels.tf`; digest `sha256:8bd9e928fbd93d358e8585ed27bd36f2f8b8b9182227ba169cfa01d646de0869`.

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

resource "xcsh_http_loadbalancer" "test" {
  name      = "example"
  namespace = "system"

  labels = {
    environment = "test"
    team        = "platform"
    managed_by  = "terraform"
  }

  domains = ["test.example.com"]

  http {
    port = 80
  }

  advertise_on_public_default_vip = {}
}
```

<a id="canonical-1021031001203102-1232322101103101-3000301221322230-3313031312223200-3223100131230111-0333323031311201-0101001312210011-1013201123131330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## With origin pool example

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md#canonical-1323101131323213-1200201313300133-0300111301103013-3131213012211311-3010002202001122-3231033222023323-2111000220211131-3013003223311203)
- [Examples](resources--http_loadbalancer--examples--group-001.md#canonical-2021030002312302-3322220033323000-0123230032002330-3323322031301112-3231001302332112-3232030323010112-1100302103320001-3132301331220203)
- With origin pool

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_http_loadbalancer/with-origin-pool.tf`; digest `sha256:108dbeb11bdb7cc69645334a91b9dce00f52f21d75be901d99e0c95036f9f959`.

```terraform
# WithOriginPool — Acceptance-test-derived Configuration
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
  port      = 443

  origin_servers {
    public_name {
      dns_name = "example.com"
    }
  }

  no_tls                = {}
  same_as_endpoint_port = {}
}

resource "xcsh_http_loadbalancer" "test" {
  name      = "example"
  namespace = "system"

  domains = ["test.example.com"]

  http {
    port = 80
  }

  default_route_pools {
    pool {
      name      = xcsh_origin_pool.test.name
      namespace = "system"
    }
    weight   = 1
    priority = 1
  }

  advertise_on_public_default_vip = {}
}
```
