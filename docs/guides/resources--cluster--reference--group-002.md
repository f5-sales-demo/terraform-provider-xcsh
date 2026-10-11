---
page_title: "xcsh_cluster reference"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cluster reference."
---

# xcsh_cluster reference

<a id="canonical-1012010313201311-3231112131122330-1332012131002120-1230002231120210-2021221111212112-1312211031310022-2223123113210212-1301210330231131"></a>

## `tls_parameters.cert_params.certificates.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3221110212020010-3222330302010201-2212232220221320-1113330020233202-0131300121301201-0201322320103000-2210031300130130-0201300232103102"></a>

<a id="canonical-0311002332002131-0312020111022032-2002031033222311-0212000211201133-0232332212210333-0202120303330203-2222230103133313-0303110131323122"></a>

## `tls_parameters.cert_params.certificates.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3230113212220110-1311312112103302-1122313301110122-2300101023102300-2311300223331201-3302320331023023-1100030020113322-2020202131030311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.cert_params.skip_server_verification` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-1100203002130122-0022303002220133-3103111323222000-2330331223110031-2130313221330022-3031012100210311-2322023011022002-1201200201010022)
- tls_parameters.cert_params.skip_server_verification

<a id="canonical-2310320120131103-3032122203113332-2012320213331310-1323110111033133-0313210321012131-1302122211302102-3322230332101313-2022233001303023"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

Terraform syntax:

```terraform
skip_server_verification = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2101011001120302-3311230102233231-2123132022111332-1330332321231133-2303000103200231-1130103300133322-3102220332233122-0200233112023111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.cert_params.tls_validation_params` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-1100203002130122-0022303002220133-3103111323222000-2330331223110031-2130313221330022-3031012100210311-2322023011022002-1201200201010022)
- tls_parameters.cert_params.tls_validation_params

<a id="canonical-1111221222123022-1311312321201030-3330332220030111-1200001323233030-3230102010132111-1113221320202030-3132103201021310-1200203223313222"></a>

Type: `"object"`. single nested block, Optional.

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

Terraform syntax:

```terraform
tls_validation_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-2132323012022112-1312002223222301-2213020020203330-0021011331133132-1320131001202203-3312133122211022-3320033323033130-1111111332330011"></a>

### Direct properties for `tls_parameters.cert_params.tls_validation_params`

<a id="canonical-0121331230300021-2131120113111110-0213302301032020-1110131321331233-1001333312301313-1312322231111333-0022130303322013-3021102311010223"></a>

#### `tls_parameters.cert_params.tls_validation_params.skip_hostname_verification` property

Type: `"bool"`. Optional.

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

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

- [trusted_ca](resources--cluster--reference--group-002.md#canonical-2021030210210031-1231023013013102-0122320101032012-2302332333111113-0333210230330213-3302100001131133-3133303313021330-3103133220131322): complete subsection reference.

<a id="canonical-0003123203233123-2103211200331331-3222100313110322-1203303132102312-1132000032100011-0232333212022332-3010100303101200-0333022303211321"></a>

<a id="canonical-1222101201320300-2000232103223202-1011310101301021-0021131322222101-1220101133103000-0100133232113223-3022110000233313-3311321333020120"></a>

#### `tls_parameters.cert_params.tls_validation_params.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-3021203102003011-2132201220102023-3103210120303002-0230212112123222-0012101202132122-0200130020330221-2221002003102323-2113110210133221"></a>

<a id="canonical-0233212032131100-0232012113022313-0300123132122130-2333010203122221-2112333020102030-0320113132222302-2121021020331210-3102121203122222"></a>

#### `tls_parameters.cert_params.tls_validation_params.verify_subject_alt_names` property

Type: `["list", "string"]`. Optional.

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

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

<a id="canonical-2021030210210031-1231023013013102-0122320101032012-2302332333111113-0333210230330213-3302100001131133-3133303313021330-3103133220131322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.cert_params.tls_validation_params.trusted_ca` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-1100203002130122-0022303002220133-3103111323222000-2330331223110031-2130313221330022-3031012100210311-2322023011022002-1201200201010022)
- [tls_parameters.cert_params.tls_validation_params](resources--cluster--reference--group-002.md#canonical-2101011001120302-3311230102233231-2123132022111332-1330332321231133-2303000103200231-1130103300133322-3102220332233122-0200233112023111)
- tls_parameters.cert_params.tls_validation_params.trusted_ca

<a id="canonical-2121302331021012-1111012132210211-1110322023000113-0323302210201011-1120212101210303-3003102332313332-3123200200121322-3223311130203221"></a>

Type: `"object"`. single nested block, Optional.

Root CA Certificate Reference. Reference to Root CA Certificate.

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-0120201212232011-1122133123331203-2212130223013020-3330131322013320-1123320202221002-3001001203111001-3321111203312233-2212223333023032"></a>

### Direct properties for `tls_parameters.cert_params.tls_validation_params.trusted_ca`

- [trusted_ca_list](resources--cluster--reference--group-002.md#canonical-3210202012233222-2222111012311100-0302200210313202-1110231121211201-2133310310320312-1123123022220111-1311011101003222-2302030122313211): complete subsection reference.

<a id="canonical-3210202012233222-2222111012311100-0302200210313202-1110231121211201-2133310310320312-1123123022220111-1311011101003222-2302030122313211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-1100203002130122-0022303002220133-3103111323222000-2330331223110031-2130313221330022-3031012100210311-2322023011022002-1201200201010022)
- [tls_parameters.cert_params.tls_validation_params](resources--cluster--reference--group-002.md#canonical-2101011001120302-3311230102233231-2123132022111332-1330332321231133-2303000103200231-1130103300133322-3102220332233122-0200233112023111)
- [tls_parameters.cert_params.tls_validation_params.trusted_ca](resources--cluster--reference--group-002.md#canonical-2021030210210031-1231023013013102-0122320101032012-2302332333111113-0333210230330213-3302100001131133-3133303313021330-3103133220131322)
- tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list

<a id="canonical-3200032031202331-3303310300201103-0322120323321011-0233001313110121-2121332331211113-2030232102122110-3323313103000220-1102031200010212"></a>

Type: `"object"`. list nested block, Optional.

Root CA Certificate Reference. Reference to Root CA Certificate.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
trusted_ca_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1133110110211033-0031313311211130-1201230132211133-2332103120313132-0010003201200300-2333012330331123-0312103322023200-3313313001323231"></a>

### Direct properties for `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list`

<a id="canonical-2302202330020102-1112111232032000-3100020330330013-3300231032212032-2122230221300202-2221120102213111-1223011102301131-3200122201101313"></a>

#### `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3112112023102223-2302020210033110-3031231230331132-3010200110110000-0222322020311310-2112313331202001-0132332120020021-3131230332331030"></a>

<a id="canonical-0131030223321101-2230030332331333-2013230103122112-2000031302300013-3023010003030323-0223213130010310-0332203213300231-3000312111332112"></a>

#### `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2003222222313001-2020000110322312-0320300121213332-2131213111123132-0111211202331003-0230110303320210-3200103230033322-1223202122310322"></a>

<a id="canonical-3210010123222102-1102310031220201-2003202223233122-3030211233221013-2003011100302121-3333132021220003-0102310211032302-1002021222303323"></a>

#### `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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

<a id="canonical-1212333233312332-3100302131201022-3112113312211001-1030232311200102-1202222120202102-2322202102313301-1202013022200232-0011002213001333"></a>

<a id="canonical-1223330312002110-2113220100032031-2102101021313213-2312312220313121-0003033020020100-3000232330300321-2303022303310300-1001133223022320"></a>

#### `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1102123322002322-1320322100011030-0010320132233300-3231000232301232-2113032101033010-2022012022113021-2230130320012023-0200202301111130"></a>

<a id="canonical-1022202212223103-2033130330220101-1312320020121032-3003202023332031-2131013231310202-0312331311012112-1003310310130030-2200233013133110"></a>

#### `tls_parameters.cert_params.tls_validation_params.trusted_ca.trusted_ca_list.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-0232001101331213-3011133133010213-0313330030310323-1311002133300112-1300011211333120-3323002231113130-0300100332301033-0311100310121120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.cert_params.volterra_trusted_ca` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.cert_params](resources--cluster--reference--group-001.md#canonical-1100203002130122-0022303002220133-3103111323222000-2330331223110031-2130313221330022-3031012100210311-2322023011022002-1201200201010022)
- tls_parameters.cert_params.volterra_trusted_ca

<a id="canonical-0313223202132222-0102310330211003-1122222300322232-0102123033120001-0033313312203210-0331233133120013-0033111120013333-2000213201212232"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for volterra trusted ca.

Additional upstream details:

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

Terraform syntax:

```terraform
volterra_trusted_ca = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- tls_parameters.common_params

<a id="canonical-0133021313320021-0231101213111210-1223331302131032-2211021112302021-0111210033212211-0312323103232222-0130100230023033-1102103031021031"></a>

Type: `"object"`. single nested block, Optional.

Information of different aspects for TLS authentication related to ciphers, certificates and trust
store.

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
common_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-3300223223231302-3231233330123300-0220102301002231-3011321013133003-1312201321331310-1321330330322113-3112330230132113-0002313303000020"></a>

### Direct properties for `tls_parameters.common_params`

<a id="canonical-1133211111101322-2021033321312003-1102022311332300-1123322233023100-0323202010300030-2103211332321002-3101113233123301-1131313200112222"></a>

#### `tls_parameters.common_params.cipher_suites` property

Type: `["list", "string"]`. Optional.

The following list specifies the supported cipher suite TLS\_AES\_128\_GCM\_SHA256
TLS\_AES\_256\_GCM\_SHA384 TLS\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_ECDHE\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_ECDHE\_RSA\_WITH\_AES\_256\_CBC\_SHA
TLS\_RSA\_WITH\_AES\_128\_CBC\_SHA TLS\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_RSA\_WITH\_AES\_256\_CBC\_SHA TLS\_RSA\_WITH\_AES\_256\_GCM\_SHA384

If not specified, the default list: TLS\_ECDHE\_ECDSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_ECDSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_RSA\_WITH\_AES\_128\_GCM\_SHA256
TLS\_ECDHE\_RSA\_WITH\_CHACHA20\_POLY1305\_SHA256 TLS\_ECDHE\_ECDSA\_WITH\_AES\_256\_GCM\_SHA384
TLS\_ECDHE\_RSA\_WITH\_AES\_256\_GCM\_SHA384 will be used.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-0320231230222310-0031001003033230-1201022023020101-2133312001211130-1231033011233132-3022001000300031-3232021111022133-0203032320021213"></a>

<a id="canonical-0220211331121010-3301123031123311-2320000022200220-0312003131103303-3101103021210031-1002020020110020-0230003201230212-1321013330033010"></a>

#### `tls_parameters.common_params.maximum_protocol_version` property

Type: `"string"`. Optional, Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2310310133130010-3323131322122011-3023012102321000-3303220210103213-2120312131130023-0200032022020211-0222120022211013-2023022123131122"></a>

<a id="canonical-0201003001300011-2220121322303203-0302023000212111-2123130100311321-1100200111110333-3203320001301322-2331110331303220-0130100212012323"></a>

#### `tls_parameters.common_params.minimum_protocol_version` property

Type: `"string"`. Optional, Computed.

\[Enum: TLS\_AUTO|TLSv1\_0|TLSv1\_1|TLSv1\_2|TLSv1\_3\] TlsProtocol is enumeration of supported TLS
versions F5 Distributed Cloud will choose the optimal TLS version. Possible values are
\`TLS\_AUTO\`, \`TLSv1\_0\`, \`TLSv1\_1\`, \`TLSv1\_2\`, \`TLSv1\_3\`. Defaults to \`TLS\_AUTO\`.

Receipt-pinned upstream constraints:

```json
{
  "default": "TLS_AUTO",
  "enum": [
    "TLS_AUTO",
    "TLSv1_0",
    "TLSv1_1",
    "TLSv1_2",
    "TLSv1_3"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

- [tls_certificates](resources--cluster--reference--group-002.md#canonical-0131010103112232-3231212333021030-2003223122303221-0310132121203233-1133132000120130-0020110303123221-1122130221010233-3121233100221002): complete subsection reference.

- [validation_params](resources--cluster--reference--group-002.md#canonical-1131330230312303-0310300220320222-1103213131021110-2122322113313130-2030303111023121-0130120220032310-3333210220201121-3331003030302103): complete subsection reference.

<a id="canonical-0131010103112232-3231212333021030-2003223122303221-0310132121203233-1133132000120130-0020110303123221-1122130221010233-3121233100221002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.common_params](resources--cluster--reference--group-002.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- tls_parameters.common_params.tls_certificates

<a id="canonical-0100100311312333-3331230103112232-3233213303223112-1233021101313233-1321110021213113-1232220323210211-2132130020232133-0023113310302132"></a>

Type: `"object"`. list nested block, Optional.

TLS Certificates. Set of TLS certificates.

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
tls_certificates {
  # Configure direct properties listed below.
}
```

<a id="canonical-1231310013231310-0222321023022133-0222300012131200-3002223100330013-0200001331233031-2011212010113321-3132330123120222-3033321230023110"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates`

- [blindfold](resources--cluster--reference--group-002.md#canonical-2201320202330010-0321030202210222-3230232012101332-2321110023300031-2032133332320203-2121203123330312-1203132001101223-0131202003122023): complete subsection reference.

<a id="canonical-3101311010203203-2023303120120233-2201030100031123-0220333013232222-0030311000132223-2203021231332020-2201023332122111-3002303203102012"></a>

<a id="canonical-1130031313212112-1112300130132321-3231303032000002-0020122233221133-1011310201013123-0331302101323130-3022021332310001-1231303123130213"></a>

#### `tls_parameters.common_params.tls_certificates.certificate_url` property

Type: `"string"`. Optional, Computed.

TLS certificate. Certificate or certificate chain in PEM format including the PEM headers.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.certificate_url": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

- [custom_hash_algorithms](resources--cluster--reference--group-002.md#canonical-3100030202202023-3332102123210012-3211221011112033-1010201322312023-3201300103121021-3032033210313123-1010332231032003-2220333022212312): complete subsection reference.

<a id="canonical-1022023000001122-3130210331230033-1122200232321113-2232212002202210-0031221200302202-1200113330312301-1102100322322223-3021301032320031"></a>

<a id="canonical-3213132120111020-0211310223200122-0331020101002031-2102012200032313-3322010301331103-2023222211023200-0021212013012133-1133231321303202"></a>

#### `tls_parameters.common_params.tls_certificates.description_spec` property

Type: `"string"`. Optional.

Description. Description for the certificate.

- [disable_ocsp_stapling](resources--cluster--reference--group-002.md#canonical-1113012131322320-0302131123233120-3101220102133203-2113133101130003-2232130312102120-0013123302000011-3200201321123120-3032320120001011): complete subsection reference.

- [private_key](resources--cluster--reference--group-002.md#canonical-0112322121010113-2232031323011101-1331110203332032-1103021020010321-1303003331100031-1100121002230200-0112002301111132-2220220321333331): complete subsection reference.

- [use_system_defaults](resources--cluster--reference--group-002.md#canonical-3312123133213002-2033110322000311-1233120302113100-3332231330331002-2223313020333233-3322023210233300-1001131001321311-3210320211002102): complete subsection reference.

<a id="canonical-2201320202330010-0321030202210222-3230232012101332-2321110023300031-2032133332320203-2121203123330312-1203132001101223-0131202003122023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.blindfold` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.common_params](resources--cluster--reference--group-002.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-002.md#canonical-0131010103112232-3231212333021030-2003223122303221-0310132121203233-1133132000120130-0020110303123221-1122130221010233-3121233100221002)
- tls_parameters.common_params.tls_certificates.blindfold

<a id="canonical-0321112000011100-1200221133001212-3300133002031023-0111133111210100-2221011012001131-0122210022310220-1110303223312031-3313233301121230"></a>

Type: `"single"`. Optional.

Native certificate preparation. Use PEM files or a P12 file, or write-only key/bundle values with
material\_version (Terraform 1.11+). Defaults to shared/ves-io-allow-volterra. Inline certificates
require unique IDs. Private inputs are never stored.

<a id="canonical-0203103211201213-2113033012000133-0012102212001323-0222123103221310-3130011020201300-2021212322022010-3301321200230331-2223112013202222"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates.blindfold`

<a id="canonical-1120022001230220-3212130023100132-2303213331012120-1013100012222102-1333133013000102-3310333023300310-2030202231232101-3131223033310011"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.algorithm` property

Type: `"string"`. Computed.

<a id="canonical-1312000320220230-0213303210222021-2123030213221313-3011301220002003-2323020001020033-1213121101232220-3101131302211200-1323011031010313"></a>

<a id="canonical-3023230300231321-2113110033003200-2232221202300233-1102023032232013-1323310112010030-1331003202021112-3132103113222121-1010030300312120"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.certificate_file` property

Type: `"string"`. Optional.

<a id="canonical-1212221010102220-2331132323301310-0020232103331212-2023211110231312-3323203010112100-3313020020110021-3231202032100321-0303222232200021"></a>

<a id="canonical-1302312011222302-0223011033030130-1011103331131302-0113123300010022-3131332111001033-1023130033131221-2001303231213133-1223033102020320"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.certificate_pem` property

Type: `"string"`. Optional.

<a id="canonical-0221231121303211-2012000202221202-3210113131122110-1313101231213032-3201302310223030-3310223130013221-3032202301012011-0133303021302232"></a>

<a id="canonical-2200110001102312-0212030033122312-2112310200330130-2211231200033000-2323022010213221-3331222110201320-1312131033012231-2111221001130201"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.chain_identity` property

Type: `"string"`. Computed.

<a id="canonical-2323322021302320-1323300121312233-1302301322112310-3321013110201311-0231232133100211-3310311112111202-0002303233102210-0033220330002231"></a>

<a id="canonical-0010332322223020-2211113130220101-3103331300302111-1032033113032323-2331333103112111-1330302020002003-1311213112133101-0332011332021302"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.context_digest` property

Type: `"string"`. Computed.

<a id="canonical-1102230010223111-2132111011323330-1200332320200100-2023210302012123-1203021331212200-3030110332001031-1300201320221022-3023012001301322"></a>

<a id="canonical-0132012203323230-3231013120322311-0021131211220232-1311231310202013-2121112321301111-1122111013320130-1203030101220221-2202223112023100"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.encrypted_location` property

Type: `"string"`. Computed, Sensitive.

<a id="canonical-3301033023311032-2112303210203231-1113100000033233-3303211331130211-2221300033112003-0120213313013013-3212010123022012-1203101030310030"></a>

<a id="canonical-1121011100031330-0110202213202023-1112031201022201-3320101023033130-3203233111232302-0003202321022332-1013032030000120-2122310223113101"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.expires_at` property

Type: `"string"`. Computed.

<a id="canonical-2102211333312201-1121211022003311-2212332201221033-2321311030133123-3303310022100330-0021131100303322-1001213311300311-2110320020323113"></a>

<a id="canonical-1121210233221230-0333221120003333-1003000323011232-0001313221333230-1132310022010123-2030333133020302-2132323101332330-1003121021322312"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.fingerprint` property

Type: `"string"`. Computed.

<a id="canonical-2302133103011311-0300002112300303-3001121203022010-3303333021311100-3321203312321003-2333320101122132-0322123230031333-2121320323133233"></a>

<a id="canonical-2021322110321230-0221300011130203-1100320033313022-1013302312321311-1030301113022133-0232310320002311-2323032320003033-0310113021113330"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.id` property

Type: `"string"`. Optional.

<a id="canonical-0103110233111132-3223321101120120-3030203000223302-1000223300231311-1323302032332123-3330030131100212-0202212333201100-0313101331301031"></a>

<a id="canonical-3032311232031022-2013010323133121-0332200323230303-0112210123113333-0110013001113230-0121101230232101-1112220023110202-0220311101111120"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.material_version` property

Type: `"string"`. Optional.

<a id="canonical-2320212313111203-2123022322132032-1031102321021001-2122013000000203-0000102212113313-2011320223030311-2222122133302203-1033000313210131"></a>

<a id="canonical-2310103322131131-2002311233313222-2220333123031210-3333103230103003-1222203111222120-3102230211322302-1023310032233131-3333031030301123"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.passphrase_env` property

Type: `"string"`. Optional.

<a id="canonical-3322023012210021-3002000313223202-2100121310033232-0201201303310200-2122101213323320-1322123221300113-3313311120100022-3001310002012011"></a>

<a id="canonical-3031310330322133-3320201022330331-2223100331213113-3110001302021211-3202031000200133-3310331000021333-3200223131101323-3011221112331021"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.passphrase_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-2303313313030312-0333130002331112-0011111102121330-3102031101132322-1330123013111131-0110110001231003-2131313233303103-1030211102213121"></a>

<a id="canonical-1211232110323002-0111003231012112-3113100303112120-3232210212022020-0023032200000120-2102330022132333-1211320212002301-0121232202130233"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.pkcs12_file` property

Type: `"string"`. Optional.

<a id="canonical-3221020220023122-3013310131032323-0321000200123303-1001321103222010-0031110233110320-2033223030303000-1321313113010013-1300212100100133"></a>

<a id="canonical-0330003003330002-1233003001033031-2113030212021200-0203103131002210-3111101131003110-3311232210123010-0131300102223302-1112222201222303"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.pkcs12_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-3302112130112133-1300030310233232-0302300023130023-1113303211211112-2232033013110322-0233231300221010-3111202222313202-3022110123223233"></a>

<a id="canonical-1133323303001201-1223101230131130-0020112213133331-2330303030023233-0033320101100202-3113022002032323-2120321231103311-0102203213023233"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.policy` property

Type: `"string"`. Optional.

<a id="canonical-3000111021210222-2001112031022230-0133033010320232-3301220333232113-1002323230203100-3110032310213103-3000313103010113-3000302223020130"></a>

<a id="canonical-3032120211113323-2013021132033300-2231211333202320-1012312123120133-1303313000310233-0010033332121003-3302033301002110-2301131213310003"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.prepared_identity` property

Type: `"string"`. Computed.

<a id="canonical-2220232321023203-2022332030323123-3022131222230112-3320122021121123-3122133322121221-0002301302222200-2101200000101101-2333033321320022"></a>

<a id="canonical-3303033003301320-0120331312030002-2211110001210120-3313233332313130-0221201013230301-2200133230210011-1112023223212211-0331233213133130"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.private_key_file` property

Type: `"string"`. Optional.

<a id="canonical-2000031120123000-3212230032222023-2213000001012022-1132311133003102-0300121312010031-1222021123203010-3110330212103111-0003302300221330"></a>

<a id="canonical-1103023203010110-2213223001301000-0321220221220030-0210123301123330-3220010100112122-1013200130330200-1131320120213110-0020102203112032"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.private_key_wo` property

Type: `"string"`. Optional, Sensitive, Write_only.

<a id="canonical-3130013211301211-0300200030002030-3100110002113101-0011211030000101-2220213113210011-0302123132012133-1111323303133010-1033031322001103"></a>

<a id="canonical-0133222200220110-3102102112211231-1002123222120220-0223211103003211-3213101211320213-0133203013300213-0213203300333332-2233102223123311"></a>

#### `tls_parameters.common_params.tls_certificates.blindfold.spki_identity` property

Type: `"string"`. Computed.

<a id="canonical-3100030202202023-3332102123210012-3211221011112033-1010201322312023-3201300103121021-3032033210313123-1010332231032003-2220333022212312"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.common_params](resources--cluster--reference--group-002.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-002.md#canonical-0131010103112232-3231212333021030-2003223122303221-0310132121203233-1133132000120130-0020110303123221-1122130221010233-3121233100221002)
- tls_parameters.common_params.tls_certificates.custom_hash_algorithms

<a id="canonical-3101031322202101-3212201230103033-2202123311010221-2101223201322020-3301002232223211-3123003233310010-1103000113333311-2031112213132123"></a>

Type: `"object"`. single nested block, Optional.

Specifies the hash algorithms to be used.

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
custom_hash_algorithms {
  # Configure direct properties listed below.
}
```

<a id="canonical-2112120330210332-0013102220002101-3111320232113031-1032333213201323-0222232233110301-1003030031321330-1031321220012122-1230323233213110"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates.custom_hash_algorithms`

<a id="canonical-1230311011232210-2013031031002200-0013231023102303-3130233203233231-3301002030231301-3200121112003311-1221311213012031-3310110210330013"></a>

#### `tls_parameters.common_params.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Optional.

\[Enum: INVALID\_HASH\_ALGORITHM|SHA256|SHA1\] Ordered list of hash algorithms to be used. Possible
values are \`INVALID\_HASH\_ALGORITHM\`, \`SHA256\`, \`SHA1\`. Defaults to
\`INVALID\_HASH\_ALGORITHM\`.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 4,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 4,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "4",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1113012131322320-0302131123233120-3101220102133203-2113133101130003-2232130312102120-0013123302000011-3200201321123120-3032320120001011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.common_params](resources--cluster--reference--group-002.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-002.md#canonical-0131010103112232-3231212333021030-2003223122303221-0310132121203233-1133132000120130-0020110303123221-1122130221010233-3121233100221002)
- tls_parameters.common_params.tls_certificates.disable_ocsp_stapling

<a id="canonical-1002301011003130-2301121301112233-3232023310303022-2202321022201023-1320221001133132-3321012311003312-1132103320123102-0131202113312312"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

Additional upstream details:

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

Terraform syntax:

```terraform
disable_ocsp_stapling = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0112322121010113-2232031323011101-1331110203332032-1103021020010321-1303003331100031-1100121002230200-0112002301111132-2220220321333331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.common_params](resources--cluster--reference--group-002.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-002.md#canonical-0131010103112232-3231212333021030-2003223122303221-0310132121203233-1133132000120130-0020110303123221-1122130221010233-3121233100221002)
- tls_parameters.common_params.tls_certificates.private_key

<a id="canonical-2200322321333323-0311000201202210-3032133223101121-3332230311011312-3013132022222033-1332222213112201-0030322122312002-0022103013103002"></a>

Type: `"object"`. single nested block, Optional.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

Terraform syntax:

```terraform
private_key {
  # Configure direct properties listed below.
}
```

<a id="canonical-0200302322003003-3200030131102030-3002300311303231-2222233012310233-1130010332332221-0223110331023232-3010122100210030-1231100112132101"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates.private_key`

- [blindfold_secret_info](resources--cluster--reference--group-002.md#canonical-1111201001212120-2100030320133102-3113220213330111-2312322131330022-1100301011312201-0210323001212333-3322202132011012-0132330022213300): complete subsection reference.

- [clear_secret_info](resources--cluster--reference--group-002.md#canonical-1003030120101221-1132033110233013-2333213312110121-0201233200102312-3032112003002113-2132331313311002-2101032023032311-3121200103302010): complete subsection reference.

<a id="canonical-1111201001212120-2100030320133102-3113220213330111-2312322131330022-1100301011312201-0210323001212333-3322202132011012-0132330022213300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.common_params](resources--cluster--reference--group-002.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-002.md#canonical-0131010103112232-3231212333021030-2003223122303221-0310132121203233-1133132000120130-0020110303123221-1122130221010233-3121233100221002)
- [tls_parameters.common_params.tls_certificates.private_key](resources--cluster--reference--group-002.md#canonical-0112322121010113-2232031323011101-1331110203332032-1103021020010321-1303003331100031-1100121002230200-0112002301111132-2220220321333331)
- tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-3311210202122101-3302221201022001-3331323203112112-2000003003132301-3310133013000021-0212222331031220-0003033221332130-3112012011231011"></a>

Type: `"object"`. single nested block, Optional.

BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.

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
blindfold_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-2021222332101311-2301232321020331-0021210213331022-1222303220102313-2310313233002300-2010012021131023-3120301330222311-2101312002323213"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-1133323020300323-2112121112330130-3002001010323102-2011303120200313-3133033102133022-3320110200201111-1020330320201221-1102110103210230"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the backend Secret
Management service.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3130211113302111-0100313000122211-2331010311133100-1313323223303123-3112111212213221-3320223223230023-1131213201230110-0300033322313312"></a>

<a id="canonical-2300313013113212-2103213031221321-2212121031101221-2223321102311022-0121121333303002-2103220200131220-0120213020130210-1333202331223213"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.location` property

Type: `"string"`. Optional, Sensitive.

Location is the URI\_ref. It could be in URL format for string:/// Or it could be a path if the
store provider is an HTTP/HTTPS location.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "content",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "maxLength": 131072,
    "metadata": {
      "category": "content",
      "confidence": 1.0,
      "note": "Blindfold envelope encryption (AES-256-GCM + RSA-OAEP) of an RSA-2048 TLS private key produces ~3700 char string:/// URL. 128KB max secret size = ~175KB base64. Discovery reported 1024 which is incorrect.",
      "source": "manual-override",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 4
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-0131002322032232-2120133203213031-3133003302323223-2133023303132200-0210221102010302-1332323333132311-2201210033133001-0233022100221313"></a>

<a id="canonical-2123320233310113-1132333310203313-0121202300330033-1003211002223220-3300000201001312-2312232011301023-2212232333210030-1120113021320030"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1003030120101221-1132033110233013-2333213312110121-0201233200102312-3032112003002113-2132331313311002-2101032023032311-3121200103302010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.common_params](resources--cluster--reference--group-002.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-002.md#canonical-0131010103112232-3231212333021030-2003223122303221-0310132121203233-1133132000120130-0020110303123221-1122130221010233-3121233100221002)
- [tls_parameters.common_params.tls_certificates.private_key](resources--cluster--reference--group-002.md#canonical-0112322121010113-2232031323011101-1331110203332032-1103021020010321-1303003331100031-1100121002230200-0112002301111132-2220220321333331)
- tls_parameters.common_params.tls_certificates.private_key.clear_secret_info

<a id="canonical-0323230330133100-2323213031330313-3021120112230220-1132220122200001-0030320320030033-2222221111120033-1331331323030133-3221303333013223"></a>

Type: `"object"`. single nested block, Optional.

ClearSecretInfoType specifies information about the Secret that is not encrypted.

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
clear_secret_info {
  # Configure direct properties listed below.
}
```

<a id="canonical-3020102121202023-0231122101031322-2331301002123111-0301010200330202-1101223321332220-1001030221021022-3030011303323113-1223022111231023"></a>

### Direct properties for `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info`

<a id="canonical-3303210333332310-3202103000302010-1112120201321221-0332122003313230-1120222123222322-2010230302000133-3132331122122123-0100022212113303"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Optional.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-1023301333020013-3301230310130222-1031102122221233-2201003100132010-2220311011000231-3030001013333223-2210031311300221-0021203211322301"></a>

<a id="canonical-1212302333232331-1002002311111223-3220321222322133-2131012222130030-2312201033113013-2300020032210220-3002221003231302-1323230103211210"></a>

#### `tls_parameters.common_params.tls_certificates.private_key.clear_secret_info.url` property

Type: `"string"`. Optional, Sensitive.

URL of the secret. Currently supported URL schemes is string:///. For string:/// scheme, Secret
needs to be encoded base64 format. When asked for this secret, caller will GET Secret bytes after
base64 decoding.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "uri",
    "formatDescription": "RFC 3986 URI with scheme (http, https, ftp)",
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^(https?|ftp)://[^\\s/$.?#].[^\\s]*$",
    "validation": {
      "rfc": "RFC 3986"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-f5xc-sensitive": true,
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.uri_ref": "true"
  }
}
```

<a id="canonical-3312123133213002-2033110322000311-1233120302113100-3332231330331002-2223313020333233-3322023210233300-1001131001321311-3210320211002102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.common_params](resources--cluster--reference--group-002.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- [tls_parameters.common_params.tls_certificates](resources--cluster--reference--group-002.md#canonical-0131010103112232-3231212333021030-2003223122303221-0310132121203233-1133132000120130-0020110303123221-1122130221010233-3121233100221002)
- tls_parameters.common_params.tls_certificates.use_system_defaults

<a id="canonical-0120003301311011-3311300110113220-3301231322330122-2331031113210112-3122223030032101-3010310323031231-3332213031013023-1103122301132320"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

Additional upstream details:

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

Terraform syntax:

```terraform
use_system_defaults = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1131330230312303-0310300220320222-1103213131021110-2122322113313130-2030303111023121-0130120220032310-3333210220201121-3331003030302103"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.validation_params` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.common_params](resources--cluster--reference--group-002.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- tls_parameters.common_params.validation_params

<a id="canonical-0133133102312101-1332332333220020-0232232312232220-0021220323221021-0021132030110220-3222222000230321-3201120003120312-1321333333221100"></a>

Type: `"object"`. single nested block, Optional.

This includes URL for a trust store, whether SAN verification is required and list of Subject Alt
Names for verification.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-trusted_ca_choice": "[\"trusted_ca\",\"trusted_ca_url\"]"
}
```

Terraform syntax:

```terraform
validation_params {
  # Configure direct properties listed below.
}
```

<a id="canonical-3113103213311210-3302310220332300-1212131110331013-2022331010002033-1211101023020002-3210002001210102-1102220233201321-3323213100111202"></a>

### Direct properties for `tls_parameters.common_params.validation_params`

<a id="canonical-2211003312223203-2030300222232211-3222310330011031-1320332023303302-0102131300203033-0232310003321230-0311211033300003-3030133333300110"></a>

#### `tls_parameters.common_params.validation_params.skip_hostname_verification` property

Type: `"bool"`. Optional.

When True, skip verification of hostname i.e. CN/Subject Alt Name of certificate is not matched to
the connecting hostname.

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

- [trusted_ca](resources--cluster--reference--group-002.md#canonical-2123020121303021-2110101300212030-3021001233002120-3111103121230120-2023100101002103-1211222000220210-2232311111312300-2313222211011230): complete subsection reference.

<a id="canonical-1133321011103332-0130311203223102-1033303111022323-3021001010013233-3220202112100111-0323130031302131-1233211222330322-3212121011322212"></a>

<a id="canonical-2232100100101230-0132032000130312-3320113230102013-2002331312220222-2323312003320020-1233301303013311-0012120001123012-3210121323333133"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca_url` property

Type: `"string"`. Optional.

Exclusive with \[trusted\_ca\] Inline Root CA Certificate.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 131072,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 131072
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 131072,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-2110310211123103-2102100003023020-1003312030333020-0213311011332002-0113213222133001-2300230301123331-2332002030300323-1102303130322300"></a>

<a id="canonical-0301122000100320-0233231121120110-0110313123203010-2120110120020223-3110000022121300-0000210033113223-2303313323330211-1323220330301330"></a>

#### `tls_parameters.common_params.validation_params.verify_subject_alt_names` property

Type: `["list", "string"]`. Optional.

List of acceptable Subject Alt Names/CN in the peer's certificate. When skip\_hostname\_verification
is false and verify\_subject\_alt\_names is empty, the hostname of the peer will be used for
matching against SAN/CN of peer's certificate.

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

<a id="canonical-2123020121303021-2110101300212030-3021001233002120-3111103121230120-2023100101002103-1211222000220210-2232311111312300-2313222211011230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.validation_params.trusted_ca` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.common_params](resources--cluster--reference--group-002.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- [tls_parameters.common_params.validation_params](resources--cluster--reference--group-002.md#canonical-1131330230312303-0310300220320222-1103213131021110-2122322113313130-2030303111023121-0130120220032310-3333210220201121-3331003030302103)
- tls_parameters.common_params.validation_params.trusted_ca

<a id="canonical-0113001301213033-0313332111203011-3221223322213132-0022130301130121-2320330012322003-3102000323101202-2002332201000203-2302030023233101"></a>

Type: `"object"`. single nested block, Optional.

Root CA Certificate Reference. Reference to Root CA Certificate.

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
trusted_ca {
  # Configure direct properties listed below.
}
```

<a id="canonical-2111030002330112-0220032202213230-2221230332110202-2121011102111112-1212031011022120-1210330103131020-0200331222010213-3003201323303000"></a>

### Direct properties for `tls_parameters.common_params.validation_params.trusted_ca`

- [trusted_ca_list](resources--cluster--reference--group-002.md#canonical-2031332030331232-1103121133023121-1011133313302011-2210130002320220-3121303200312203-2233202100102010-1012212230111202-0333000000121000): complete subsection reference.

<a id="canonical-2031332030331232-1103121133023121-1011133313302011-2210130002320220-3121303200312203-2233202100102010-1012212230111202-0333000000121000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- [tls_parameters.common_params](resources--cluster--reference--group-002.md#canonical-3000332102231031-2232112022123031-0211133232202112-0032122020002033-2212122300111002-2302023030003012-0222333220103023-0320013312123031)
- [tls_parameters.common_params.validation_params](resources--cluster--reference--group-002.md#canonical-1131330230312303-0310300220320222-1103213131021110-2122322113313130-2030303111023121-0130120220032310-3333210220201121-3331003030302103)
- [tls_parameters.common_params.validation_params.trusted_ca](resources--cluster--reference--group-002.md#canonical-2123020121303021-2110101300212030-3021001233002120-3111103121230120-2023100101002103-1211222000220210-2232311111312300-2313222211011230)
- tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list

<a id="canonical-1313122211012132-2323103001001011-0212010311201302-1221103302011231-3213123332332012-1223113322323302-0103033010333230-2311102132201233"></a>

Type: `"object"`. list nested block, Optional.

Root CA Certificate Reference. Reference to Root CA Certificate.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "1"
  }
}
```

Terraform syntax:

```terraform
trusted_ca_list {
  # Configure direct properties listed below.
}
```

<a id="canonical-1110332230020023-3122002023222023-1303001331113302-3130001320100030-0032200103100332-2110023113303022-0322110210011212-1112200200021323"></a>

### Direct properties for `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list`

<a id="canonical-3012103321133010-1133133303032000-1031310302000021-1311033132220110-1221023310030311-2002232310120001-3320023323313130-2101322023022232"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.kind` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then kind will hold the
referred object's kind (e.g. 'route').

Additional upstream details:

Virtual\_host) refers to another(e.g route) then kind will hold the referred object's kind (e.g.
"route")

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-2101333123213122-0023033112202223-1013022330210321-0102133302312312-0201223303312123-3012032220321220-0232010010213310-2302203210300010"></a>

<a id="canonical-0121103323201301-0311211010312120-0130311023323321-0213210203211313-1232003320112010-0302201032121132-2002020132030010-0311230330312212"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.name` property

Type: `"string"`. Optional.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-1000103232230201-0210021202300100-3030102033020013-0331201222231322-2311122312303331-3202022113301330-0011120131201310-2032023101211310"></a>

<a id="canonical-3032221332033021-0031111312101212-0302233032032031-0131131013320003-1311330220221312-2203220103321131-2302213322132223-1022011213202123"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.namespace` property

Type: `"string"`. Optional, Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "naming",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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

<a id="canonical-0122102332323321-1130221101222111-3212312033312222-2002302321330120-3101201003302113-3033232002312110-1320322320030312-0122210312110213"></a>

<a id="canonical-1231100121120022-3223203313310023-1000131301103312-2021001222321201-1111332131300100-3102212333302121-2120300103223001-1012123020200301"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3303133130220113-1232100232231223-3023230300311322-3122002310312030-1310232312133203-0031002020131301-3222202323233310-0333312110012122"></a>

<a id="canonical-0330203331233030-2232313131123310-1103310130210321-1132030000112110-0300033123203102-0020022002312033-2333130123330002-1220111330031333"></a>

#### `tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list.uid` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then uid will hold the
referred object's(e.g. Route's) uid.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
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

<a id="canonical-3111210320132103-3023301033212102-3010212200103223-1330312002311102-2203031033111312-1310030010122203-0223131200231111-2130111222200133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.default_session_key_caching` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- tls_parameters.default_session_key_caching

<a id="canonical-0303022213110012-1300303123201131-3303103203230323-0022123313320131-0210210011130022-3223230022013201-1311321323330233-1010212312123302"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default session key caching.

Additional upstream details:

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

Terraform syntax:

```terraform
default_session_key_caching = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102323013232123-2122130023210330-2223011133213113-3010120223212310-3033220231110101-2003103300202020-3311031203312302-0210012033113310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.disable_session_key_caching` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- tls_parameters.disable_session_key_caching

<a id="canonical-2303301023023213-3311333113121301-1002223201031203-1332130331003033-0322113131301333-3013101131220232-3233111300232102-0233122312021011"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable session key caching.

Additional upstream details:

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

Terraform syntax:

```terraform
disable_session_key_caching = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1231201030310032-2232133130332230-3102203310013132-0231112133233202-2112002222030130-0312121223003313-1302023203330231-0310011232322132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.disable_sni` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- tls_parameters.disable_sni

<a id="canonical-3011031101302320-2132321032112032-3201030210032322-0003233113001212-0133230213031010-0031313020011233-0001000333100302-2003230021320213"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable sni.

Additional upstream details:

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

Terraform syntax:

```terraform
disable_sni = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3303030113320302-0313223030232122-1230002312320323-0112201033032021-3032303300003222-3030123202200331-3032022101232333-0330331332030030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `tls_parameters.use_host_header_as_sni` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [tls_parameters](resources--cluster--reference--group-001.md#canonical-3220210122031322-3121113012222331-0213030232301003-2002100301202230-1030100031230032-2222122001310100-2022213213020103-1031100333310031)
- tls_parameters.use_host_header_as_sni

<a id="canonical-2012212221130301-0303332301113230-3112332122220113-2220031222031121-0110221211330111-3111022033002232-2012200203321233-2010130310233323"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

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

Terraform syntax:

```terraform
use_host_header_as_sni = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2112130020110230-3121122220021332-0102101233321233-3011130330023132-3123211112301313-2313230113011333-2220110203131001-2022033121110202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upstream_conn_pool_reuse_type` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- upstream_conn_pool_reuse_type

<a id="canonical-2123332212301000-2200010021323223-2330213113021121-0221021110000130-0001101030103313-0133101111321103-0322301331033113-0121000000212220"></a>

Type: `"object"`. single nested block, Optional.

Select upstream connection pool reuse state for every downstream connection. This configuration
choice is for HTTP(S) LB only.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-map_downstream_to_upstream_conn_pool_type": "[\"disable_conn_pool_reuse\",\"enable_conn_pool_reuse\"]"
}
```

Terraform syntax:

```terraform
upstream_conn_pool_reuse_type {
  # Configure direct properties listed below.
}
```

<a id="canonical-3320331001230102-2231121201220333-1110302213302303-0233320111111023-3233022310230302-2100003101020110-0131021233211200-1103130213111101"></a>

### Direct properties for `upstream_conn_pool_reuse_type`

- [disable_conn_pool_reuse](resources--cluster--reference--group-002.md#canonical-3323023220002122-1301220010310231-1031013200230113-2100310231030123-2113121100122000-2033203013032233-0220222232110011-1222231110321022): complete subsection reference.

- [enable_conn_pool_reuse](resources--cluster--reference--group-002.md#canonical-2303230031132101-0221312032202231-3131333202122000-0321210220013310-2012022020012310-2203211202111212-1223313210332112-0131300233211203): complete subsection reference.

<a id="canonical-3323023220002122-1301220010310231-1031013200230113-2100310231030123-2113121100122000-2033203013032233-0220222232110011-1222231110321022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upstream_conn_pool_reuse_type.disable_conn_pool_reuse` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [upstream_conn_pool_reuse_type](resources--cluster--reference--group-002.md#canonical-2112130020110230-3121122220021332-0102101233321233-3011130330023132-3123211112301313-2313230113011333-2220110203131001-2022033121110202)
- upstream_conn_pool_reuse_type.disable_conn_pool_reuse

<a id="canonical-0313323233003322-0302021233211011-3111230310312230-3103300120210113-0212200233332101-0102223123320220-0331313020002320-3113233303110330"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable conn pool reuse.

Additional upstream details:

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

Terraform syntax:

```terraform
disable_conn_pool_reuse = {}
```

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2303230031132101-0221312032202231-3131333202122000-0321210220013310-2012022020012310-2203211202111212-1223313210332112-0131300233211203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upstream_conn_pool_reuse_type.enable_conn_pool_reuse` properties

Breadcrumbs:

- [xcsh_cluster](../resources/cluster.md#canonical-0200232312311331-3030233031121222-1103023001033003-2132323312102233-3022210011323301-1323303310030333-2031301002313333-2332312221310003)
- [Property reference](resources--cluster--reference--group-001.md#canonical-3101131133022110-3312300202101321-1032132322320210-3102023201100003-0111113110022121-3023312323003101-2100300213113301-3113003202303321)
- [upstream_conn_pool_reuse_type](resources--cluster--reference--group-002.md#canonical-2112130020110230-3121122220021332-0102101233321233-3011130330023132-3123211112301313-2313230113011333-2220110203131001-2022033121110202)
- upstream_conn_pool_reuse_type.enable_conn_pool_reuse

<a id="canonical-2121321022333202-2233112003101001-3233223201112031-1201321001032030-2020012003332033-3313010300213130-2123000233333301-0313003330003133"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable conn pool reuse.

Additional upstream details:

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

Terraform syntax:

```terraform
enable_conn_pool_reuse = {}
```

This is an empty object or choice marker. It has no direct properties.
