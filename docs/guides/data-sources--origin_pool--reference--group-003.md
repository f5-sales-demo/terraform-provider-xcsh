---
page_title: "xcsh_origin_pool reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool reference."
---

# xcsh_origin_pool reference

<a id="canonical-3312011321001100-3320123303202211-1201223331112313-2021131122222212-2020103133123202-0211110031310122-0312222113202112-1113223312110123"></a>

## `origin_servers.k8s_service.site_locator.virtual_site.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1030131222233221-0131322020130322-1013021132020111-2133202323001322-1101221232121133-2021213021212031-2033130130231003-0232303312300202"></a>

<a id="canonical-0002333032300210-2321320223002001-2001223110232333-1023213033132220-1111003112312230-1030131002222211-2130213000103331-3212333012213113"></a>

## `origin_servers.k8s_service.site_locator.virtual_site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1003201010222201-1101110320120012-1210031020121301-2311123310120002-3313110220322012-2200010133303101-3120221211111312-0310000310102010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- origin_servers.k8s_service.snat_pool

<a id="canonical-2300132112011133-1001222102231111-1021033011211330-1013213202333130-1031330203222122-1031221000101110-0232123121022112-2111300211032002"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

<a id="canonical-2302000213203122-2322302201010312-3131020110113132-0112030121330102-3210120200001000-0310023231022100-1021223112320332-2233002030300011"></a>

### Direct properties for `origin_servers.k8s_service.snat_pool`

- [no_snat_pool](data-sources--origin_pool--reference--group-003.md#canonical-2021331300301000-0201112021010013-3030120000311113-3213033221203332-1322203033021022-0000202221101311-2200013212221201-1311010002330321): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-003.md#canonical-1331300313220011-3023330120132021-3001001101013210-3130003312112001-1122121110300022-3203003032232113-2301002111100331-2233032032111230): complete subsection reference.

<a id="canonical-2021331300301000-0201112021010013-3030120000311113-3213033221203332-1322203033021022-0000202221101311-2200013212221201-1311010002330321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- [origin_servers.k8s_service.snat_pool](data-sources--origin_pool--reference--group-003.md#canonical-1003201010222201-1101110320120012-1210031020121301-2311123310120002-3313110220322012-2200010133303101-3120221211111312-0310000310102010)
- origin_servers.k8s_service.snat_pool.no_snat_pool

<a id="canonical-2112133302313122-0101020213222322-2101220231122310-1330323112021031-2030200331033020-1031013313100001-2013232120021223-0320212020313322"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no snat pool.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1331300313220011-3023330120132021-3001001101013210-3130003312112001-1122121110300022-3203003032232113-2301002111100331-2233032032111230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.snat_pool.snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- [origin_servers.k8s_service.snat_pool](data-sources--origin_pool--reference--group-003.md#canonical-1003201010222201-1101110320120012-1210031020121301-2311123310120002-3313110220322012-2200010133303101-3120221211111312-0310000310102010)
- origin_servers.k8s_service.snat_pool.snat_pool

<a id="canonical-0301323033010103-2011122000221231-0203013123323002-2111221332012030-2323013333123330-0303301231023202-2322322230201030-1122122000331310"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-2102020330201322-0102303331123001-3331100200123323-3020311003210332-0122212213132133-2121123231330003-2233302021000231-2113231030122221"></a>

### Direct properties for `origin_servers.k8s_service.snat_pool.snat_pool`

<a id="canonical-1200301233210100-3313112203023312-2100103121232311-2202010012221101-2330320031030010-3031220212201302-3132221301132002-1030022001210021"></a>

#### `origin_servers.k8s_service.snat_pool.snat_pool.prefixes` property

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3003320220101232-2222122002000011-0122233212223100-0113213320301211-3321322300223222-1233131112321100-3101110202223112-2221113000223003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.k8s_service.vk8s_networks` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- origin_servers.k8s_service.vk8s_networks

<a id="canonical-1121320221330002-3130133132310313-3123111112030121-2200020223033312-3302202200233301-1203100203000022-2113333233111312-0022330210032111"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for vk8s networks.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- origin_servers.private_ip

<a id="canonical-1323132032102301-0213302210220300-1020121233103023-1020311222111112-1021313101312113-1100120021102132-1233223302202231-0132123313103321"></a>

Type: `"single"`. Computed.

Specify origin server with private or public IP address and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]",
  "x-ves-oneof-field-private_ip_choice": "[\"ip\"]"
}
```

<a id="canonical-0131000333221033-2223311223311001-0302133323312210-1333030210203222-2231323330022221-1030022220213331-1200112100100032-3333203022330310"></a>

### Direct properties for `origin_servers.private_ip`

- [inside_network](data-sources--origin_pool--reference--group-003.md#canonical-1023322133300312-3212322320031221-3000023221010020-2322001122012021-2011020320000233-1313013221012320-2120101222312123-0030101330210022): complete subsection reference.

<a id="canonical-0320121111021001-2203023121330110-3022120020120210-2221003230002312-3113001022013212-2213122121300300-1322013112110201-3332221331003322"></a>

<a id="canonical-3201122320201100-2232003003132312-1302121322103210-0213323033333222-3121122113032011-1020130230302311-3223031113311000-0333322103003223"></a>

#### `origin_servers.private_ip.ip` property

Type: `"string"`. Computed.

IP. Exclusive with \[\] Private IPv4 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [outside_network](data-sources--origin_pool--reference--group-003.md#canonical-0302202131230001-0233123210023202-0133201202333100-2322210001221331-3212233002323211-0121012333222123-2201123232222110-3213201123133132): complete subsection reference.

- [segment](data-sources--origin_pool--reference--group-003.md#canonical-3132303123311030-2330310310330102-0101200222301031-3110233301111120-0233230013023010-0223121320310230-2112211212011300-1103020210122222): complete subsection reference.

- [site_locator](data-sources--origin_pool--reference--group-003.md#canonical-1230220321001223-0010023132100120-0123100223321110-3121021330010013-0322133233101110-1121023020112003-2120101221002232-1022030000212030): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-003.md#canonical-0333322312323021-0312203002230202-3222301233233110-0330230100212001-1311203121121230-1212210203320003-3133230201312303-3231031023313001): complete subsection reference.

<a id="canonical-1023322133300312-3212322320031221-3000023221010020-2322001122012021-2011020320000233-1313013221012320-2120101222312123-0030101330210022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.inside_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-003.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- origin_servers.private_ip.inside_network

<a id="canonical-1332333132310022-3003132113211323-0322120123232001-1331322020021023-3322100031003020-3013333211202122-3030113233110100-1321000310023030"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside network.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0302202131230001-0233123210023202-0133201202333100-2322210001221331-3212233002323211-0121012333222123-2201123232222110-3213201123133132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.outside_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-003.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- origin_servers.private_ip.outside_network

<a id="canonical-3031033110331132-3203320200310102-3330333312113313-0203111020203111-0102332121203313-0001011113021332-3132111312320230-3333102001033230"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside network.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3132303123311030-2330310310330102-0101200222301031-3110233301111120-0233230013023010-0223121320310230-2112211212011300-1103020210122222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.segment` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-003.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- origin_servers.private_ip.segment

<a id="canonical-3201111022301122-3212230121302012-1222322302333010-2301231110023000-2310203230000220-3201031222233121-1110110313332120-2123331211011100"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-0130113012121210-2201002231312211-2220002010223302-1020300213230230-3200101330111112-3230021223113103-1033133033230120-0013123201302031"></a>

### Direct properties for `origin_servers.private_ip.segment`

<a id="canonical-1121332010322010-0331001213013201-2131310002211300-2222320222312010-2102333233133131-1130310232202313-0022013011121321-0310220023123313"></a>

#### `origin_servers.private_ip.segment.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3302312100203322-3331010132203231-1003021303133023-0020113022020203-3011220013333200-2233100110332001-1330202010303030-0202121321332011"></a>

<a id="canonical-1231301232000220-3333303332321031-0002022301210111-1220221031031000-0313110233211110-3330222013200311-1001023222032333-2112102111012020"></a>

#### `origin_servers.private_ip.segment.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3233332010001012-1110311120001213-3120222011223023-0333222232121211-0000220033100120-0221202320211231-2333003012020310-1123022233113212"></a>

<a id="canonical-2023330212202102-3122031020130003-2211100102101231-3111223012221202-3321321130123233-2220310031322110-3000002023123112-1111311330233122"></a>

#### `origin_servers.private_ip.segment.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1230220321001223-0010023132100120-0123100223321110-3121021330010013-0322133233101110-1121023020112003-2120101221002232-1022030000212030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.site_locator` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-003.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- origin_servers.private_ip.site_locator

<a id="canonical-1122123131011031-0313203023211013-3301232131102222-0032133132333001-3213212231100211-0032200303000013-2221221320000103-0130032103113303"></a>

Type: `"single"`. Computed.

This message defines a reference to a site or virtual site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-2101123010023032-1213333020121123-2032022222030201-1131113333303023-2111010111011011-3002313131223221-2313003012102120-2132120110112203"></a>

### Direct properties for `origin_servers.private_ip.site_locator`

- [site](data-sources--origin_pool--reference--group-003.md#canonical-1222313013302012-2300030200323111-3003211100300122-2203322122311323-2300302211112100-0001033202331322-2133203310023002-2101221311103322): complete subsection reference.

- [virtual_site](data-sources--origin_pool--reference--group-003.md#canonical-3010012311121222-3323321022012103-2213010123123311-0101103202010312-1332101333222321-3332100203330302-1302312132133132-1230033300301223): complete subsection reference.

<a id="canonical-1222313013302012-2300030200323111-3003211100300122-2203322122311323-2300302211112100-0001033202331322-2133203310023002-2101221311103322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.site_locator.site` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-003.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- [origin_servers.private_ip.site_locator](data-sources--origin_pool--reference--group-003.md#canonical-1230220321001223-0010023132100120-0123100223321110-3121021330010013-0322133233101110-1121023020112003-2120101221002232-1022030000212030)
- origin_servers.private_ip.site_locator.site

<a id="canonical-0210110200112220-0223300031210231-3200011201331330-0013221232313333-3331130210310133-3320221312023331-0301221002100220-1112123322000300"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-2113032330203103-3022033332230220-3033321231332130-0203232030131330-2203231012123220-2123132131101131-0201132122113320-2122010330132100"></a>

### Direct properties for `origin_servers.private_ip.site_locator.site`

<a id="canonical-2011333133310033-3111112013320011-3311200323000133-0010313130202330-2012112031033000-1110111312320022-1222221033022123-1112023023123102"></a>

#### `origin_servers.private_ip.site_locator.site.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1222002321100201-1101201301133331-1032323133222112-0213211133312233-0022003212323222-3312033300302221-3000210221322310-0302320302222200"></a>

<a id="canonical-1033210032023010-2201313323020200-3021320220023021-2222323113212230-3212131111323110-3310133023231221-0130111203313302-2000212221222111"></a>

#### `origin_servers.private_ip.site_locator.site.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0012123211112011-3230103211311313-3213113312212130-1022200021003102-1320213220020100-0332331223203301-2221122022102330-2203021222331130"></a>

<a id="canonical-1031201100021313-1202213331112132-2103112211211102-2321330212202320-3031231202201210-1011132120123021-2212021313231022-2130100003020233"></a>

#### `origin_servers.private_ip.site_locator.site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3010012311121222-3323321022012103-2213010123123311-0101103202010312-1332101333222321-3332100203330302-1302312132133132-1230033300301223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-003.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- [origin_servers.private_ip.site_locator](data-sources--origin_pool--reference--group-003.md#canonical-1230220321001223-0010023132100120-0123100223321110-3121021330010013-0322133233101110-1121023020112003-2120101221002232-1022030000212030)
- origin_servers.private_ip.site_locator.virtual_site

<a id="canonical-2322223312023130-3220100111322322-2331222033031323-2020102013312300-1021033311223033-1331102312103133-0121121111322321-0311300313201102"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-2310310230102303-3101320132323313-0321002031213213-0322100203331203-2103123313000300-2230101210131221-0031212111211310-3010203103212221"></a>

### Direct properties for `origin_servers.private_ip.site_locator.virtual_site`

<a id="canonical-2122220122323033-3011313232113110-1312231103032221-3021110310220201-0203022032302212-0313101011012210-3311303333111112-2111311031003330"></a>

#### `origin_servers.private_ip.site_locator.virtual_site.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3010330231003330-1332110232022312-1132231023111310-0013013003321012-1101130010312333-3221001211232020-0131332221213310-1300220133211200"></a>

<a id="canonical-2222301011330011-2000300223302330-1133220231301120-2222111311223132-0313131312122012-1320203102321033-0232323233032322-0022211231013021"></a>

#### `origin_servers.private_ip.site_locator.virtual_site.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0130000200222002-0203230011223020-0103030211123323-3122302010302333-3200032002020011-2032123202321320-1311320230023023-1120103102102012"></a>

<a id="canonical-2310231220122010-0003202132001330-0002203212112212-0132120021312021-2332302030030021-2011030033100232-2013211110113320-3323002022220322"></a>

#### `origin_servers.private_ip.site_locator.virtual_site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0333322312323021-0312203002230202-3222301233233110-0330230100212001-1311203121121230-1212210203320003-3133230201312303-3231031023313001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-003.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- origin_servers.private_ip.snat_pool

<a id="canonical-1230021313222323-3120023031131200-1233113301020201-2232222010011223-0032021312332211-3100132312300121-2202330322301313-1303211031130322"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

<a id="canonical-1202000320022012-0221120302312322-3020211100110111-2212210123303210-2221123221121310-3020212112231020-0322301111302311-3023330232221113"></a>

### Direct properties for `origin_servers.private_ip.snat_pool`

- [no_snat_pool](data-sources--origin_pool--reference--group-003.md#canonical-0232130210221013-1312312111333323-1112123110200123-0221202303332121-0121201232331122-0013121023131022-3201122223222003-1202103113231310): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-003.md#canonical-3021323101120213-2322100033001030-2033332011201330-3221222010123210-2102132032131311-0232023301002122-2021123311201013-3032110201211220): complete subsection reference.

<a id="canonical-0232130210221013-1312312111333323-1112123110200123-0221202303332121-0121201232331122-0013121023131022-3201122223222003-1202103113231310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-003.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- [origin_servers.private_ip.snat_pool](data-sources--origin_pool--reference--group-003.md#canonical-0333322312323021-0312203002230202-3222301233233110-0330230100212001-1311203121121230-1212210203320003-3133230201312303-3231031023313001)
- origin_servers.private_ip.snat_pool.no_snat_pool

<a id="canonical-0003302321133221-3130232212332103-1122210001002331-0010132303311003-1020030123022012-0012330002212132-2232113131203303-2021031103210122"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no snat pool.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3021323101120213-2322100033001030-2033332011201330-3221222010123210-2102132032131311-0232023301002122-2021123311201013-3032110201211220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_ip.snat_pool.snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-003.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- [origin_servers.private_ip.snat_pool](data-sources--origin_pool--reference--group-003.md#canonical-0333322312323021-0312203002230202-3222301233233110-0330230100212001-1311203121121230-1212210203320003-3133230201312303-3231031023313001)
- origin_servers.private_ip.snat_pool.snat_pool

<a id="canonical-2232333332130303-0330120201011332-0022211113221003-2332200023030010-0101132031231212-3132020113213211-2110030323233031-0133030103013230"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-3202103031000303-3302213232000101-1211112202321223-3010210223331310-1013231330301113-1030300011201323-2021320200223220-3120222330002311"></a>

### Direct properties for `origin_servers.private_ip.snat_pool.snat_pool`

<a id="canonical-2223130122312001-3221130033133232-3002313333331210-2202010202231213-0300330220003120-2310133320022001-1202230123122322-2301032302202013"></a>

#### `origin_servers.private_ip.snat_pool.snat_pool.prefixes` property

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- origin_servers.private_name

<a id="canonical-1212131321023331-3323221213100213-3113030213130331-1022001321131223-0300113331232220-0221212323102122-0132231303200201-0101122223310000"></a>

Type: `"single"`. Computed.

Specify origin server with private or public DNS name and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"segment\"]"
}
```

<a id="canonical-0203000131123132-0112111212320010-1213330202100103-1113223131013332-1210123233211322-1200121001001021-1321121033223022-1230010101012010"></a>

### Direct properties for `origin_servers.private_name`

<a id="canonical-0033110221311322-0232002212203030-0222030202220121-0233103311122200-2321020110302131-0120212321323311-3003313013211121-1110232232323213"></a>

#### `origin_servers.private_name.dns_name` property

Type: `"string"`. Computed.

DNS Name. DNS Name

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [inside_network](data-sources--origin_pool--reference--group-003.md#canonical-2201113023111301-2130033123212030-3101303112332102-0120230011113310-2333311132231210-3112223303322003-1233102300113011-1012011010000313): complete subsection reference.

- [outside_network](data-sources--origin_pool--reference--group-003.md#canonical-2333001322321030-1223130232211212-2231131122201102-2330212123111332-0233011020220031-1020133300202303-1231200222232112-1121230103123000): complete subsection reference.

<a id="canonical-0010301302220223-0122112222213313-1312231113030322-0333132012223023-3011222311030123-0201203210113311-1022202231103030-2102203101220321"></a>

<a id="canonical-1131123131021030-1032110313013232-2220223221223213-0231330131110222-0312312020203132-0113202232013031-1212213323323233-1123102222223031"></a>

#### `origin_servers.private_name.refresh_interval` property

Type: `"number"`. Computed.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

- [segment](data-sources--origin_pool--reference--group-003.md#canonical-2102212203301223-1333332333313011-3203313003022320-3220031002101131-2022031112111233-0123121103301132-3023002031131212-3100331232111010): complete subsection reference.

- [site_locator](data-sources--origin_pool--reference--group-003.md#canonical-2110032230130213-1233131020100213-0103031331322233-3132300222023320-1221130110230312-3122233023110122-0213032231121113-3000120212301203): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-003.md#canonical-3312011332131233-1230133131313010-3013003121231323-0223212213120012-0301213303101202-1000312320032330-3210122031212113-2000030222032011): complete subsection reference.

<a id="canonical-2201113023111301-2130033123212030-3101303112332102-0120230011113310-2333311132231210-3112223303322003-1233102300113011-1012011010000313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.inside_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-003.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- origin_servers.private_name.inside_network

<a id="canonical-3312013330201201-3132311111230322-2232223022212333-1013203331200012-3033300123203022-0000032033020033-2030333301231201-1330312212233300"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside network.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2333001322321030-1223130232211212-2231131122201102-2330212123111332-0233011020220031-1020133300202303-1231200222232112-1121230103123000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.outside_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-003.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- origin_servers.private_name.outside_network

<a id="canonical-0113310102130113-3100201313013101-2030200220102110-3133110031312120-0122202002021032-1033001331120103-3012332133022000-2321311332323123"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside network.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2102212203301223-1333332333313011-3203313003022320-3220031002101131-2022031112111233-0123121103301132-3023002031131212-3100331232111010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.segment` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-003.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- origin_servers.private_name.segment

<a id="canonical-1221001011021022-0033011011130002-2021222002011013-3120212333000223-1003102102031003-2130213321103222-2300300011213110-0222103112022322"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-2202103222301300-3312332302330022-0322223201231222-1201130123003123-0022322012003012-3013323033211213-1022331330100132-1333301302300122"></a>

### Direct properties for `origin_servers.private_name.segment`

<a id="canonical-1033202133100110-1322011323001330-3121210233103123-3000223222300302-0022212132321211-3301003002033303-0331130131221200-2100313113312010"></a>

#### `origin_servers.private_name.segment.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-2321123322120313-0312230231131300-3203012323333301-0203313112112332-1323113320100331-2022010103130201-2210113323133033-0012221331203010"></a>

<a id="canonical-2313001333100322-2012011123010210-1023221332013302-1103113311302232-1300123032021222-0013332102032033-1313211203003011-0120033331032212"></a>

#### `origin_servers.private_name.segment.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1210202331220112-1210321233303002-3021313223331210-2311030032232011-0032010311333002-0010132130313133-2122303313121013-1213230132101223"></a>

<a id="canonical-0223322313332132-0023230313222001-0000223330212000-2232102203132302-3111030212201211-1103312100133021-3010203112003200-2220203330002211"></a>

#### `origin_servers.private_name.segment.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2110032230130213-1233131020100213-0103031331322233-3132300222023320-1221130110230312-3122233023110122-0213032231121113-3000120212301203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.site_locator` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-003.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- origin_servers.private_name.site_locator

<a id="canonical-1323033030323301-1000213232102213-1222113200010031-2233202002201211-0312021122012220-3011323321322103-1222233021120013-3212331312020123"></a>

Type: `"single"`. Computed.

This message defines a reference to a site or virtual site object.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"site\",\"virtual_site\"]"
}
```

<a id="canonical-0102322022211200-1033211210333212-1301021012330101-2112210310110111-3233221111001320-0202302331130020-0331332113203201-0022222200122333"></a>

### Direct properties for `origin_servers.private_name.site_locator`

- [site](data-sources--origin_pool--reference--group-003.md#canonical-0032313212303101-0102131210033022-3013302311001130-1212201303121120-0132030031100121-1301121323310103-0023103133011303-1013221113323200): complete subsection reference.

- [virtual_site](data-sources--origin_pool--reference--group-003.md#canonical-0133303123032010-2303131222113200-2231102200221213-2211200102122101-3023311211113031-0011312030031003-1302130011213010-3030330302200322): complete subsection reference.

<a id="canonical-0032313212303101-0102131210033022-3013302311001130-1212201303121120-0132030031100121-1301121323310103-0023103133011303-1013221113323200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.site_locator.site` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-003.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- [origin_servers.private_name.site_locator](data-sources--origin_pool--reference--group-003.md#canonical-2110032230130213-1233131020100213-0103031331322233-3132300222023320-1221130110230312-3122233023110122-0213032231121113-3000120212301203)
- origin_servers.private_name.site_locator.site

<a id="canonical-2122103123032102-2131333320013322-0213003013010301-3310201212001002-3211000232030113-1101112032102111-0331323101012220-3003001010121201"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-0303131330212110-0311233033100022-3132002020000232-0122023221122003-3112021332000321-3301332020010103-0203103300133301-2202013112111311"></a>

### Direct properties for `origin_servers.private_name.site_locator.site`

<a id="canonical-1211233032113220-1001202021313232-0311202130333311-2311300302311102-2231311323110320-0001210120300302-1013101012110102-3130110300020010"></a>

#### `origin_servers.private_name.site_locator.site.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1120032123131211-0002033102121322-0331300123112000-1313233111010033-3111132113013230-3102001012122031-1120231131332003-3110312321231120"></a>

<a id="canonical-2200203231331210-2032002320223211-2201023031313130-1130232221331302-1203200300201023-3300230220002312-1001003212112330-0101112102321102"></a>

#### `origin_servers.private_name.site_locator.site.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1233203330101211-3310220211203310-1331130003213323-2010021330032230-3012231023201013-0111020013232323-3231231320110022-3131321310111031"></a>

<a id="canonical-0131020200201311-0030300021011221-0303122232320102-0101111302223020-2131121023233331-0232330221031312-2313011321300132-3133130301123111"></a>

#### `origin_servers.private_name.site_locator.site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0133303123032010-2303131222113200-2231102200221213-2211200102122101-3023311211113031-0011312030031003-1302130011213010-3030330302200322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.site_locator.virtual_site` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-003.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- [origin_servers.private_name.site_locator](data-sources--origin_pool--reference--group-003.md#canonical-2110032230130213-1233131020100213-0103031331322233-3132300222023320-1221130110230312-3122233023110122-0213032231121113-3000120212301203)
- origin_servers.private_name.site_locator.virtual_site

<a id="canonical-1030132131011010-1212123123013211-0221012011230212-1223121332332322-1032232202022121-3011331310213102-0103033321011323-0012301013230133"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-0130133110323230-1023312321230323-2021021322011030-3020012001122123-2101112220301323-3312031023103023-3331120030332222-1021233303113002"></a>

### Direct properties for `origin_servers.private_name.site_locator.virtual_site`

<a id="canonical-3110033030101323-3233302100100230-0200301100201200-0131232232012032-2230100221120101-3122323311002221-0222300103032322-0222100011303122"></a>

#### `origin_servers.private_name.site_locator.virtual_site.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-0333221002022120-0210210333101213-2011102100013021-3230220323022303-1111212231133021-0232103021320203-0231313231220120-0333233121230300"></a>

<a id="canonical-0330321331222110-0231313120110032-1030211030232202-1132331301300303-2112330332220322-0202311120001311-3112213233031103-2131213322010103"></a>

#### `origin_servers.private_name.site_locator.virtual_site.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1200122031023023-1331310113102020-1301021313002211-1130213000132332-1322121113230110-2320102132233220-3013232220001133-0232313131000023"></a>

<a id="canonical-1021020023123132-1120320000223033-1012301102030330-2021013331102200-3213232323122200-0321023330230312-3210030111210202-3213131132013103"></a>

#### `origin_servers.private_name.site_locator.virtual_site.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3312011332131233-1230133131313010-3013003121231323-0223212213120012-0301213303101202-1000312320032330-3210122031212113-2000030222032011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-003.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- origin_servers.private_name.snat_pool

<a id="canonical-1331233212211320-2330123312023213-0133103211121212-0120101132022101-0030332333002120-1303321100121230-3122021023133032-1301312121300213"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-snat_pool_choice": "[\"no_snat_pool\",\"snat_pool\"]"
}
```

<a id="canonical-1001001321300102-0012323120021223-1221133112013130-0311003232311112-0311313101023312-0310213022300313-2031223332122010-3221301330321010"></a>

### Direct properties for `origin_servers.private_name.snat_pool`

- [no_snat_pool](data-sources--origin_pool--reference--group-003.md#canonical-3312120121220023-3230131203113023-1222013132133331-2010313033233230-2332333210322111-2102020101101330-2111331013210333-1133201121322223): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-003.md#canonical-3212302120120200-1323200000301103-0223232010213033-3103221203300131-0131233211123212-0100110203310302-1021032313011123-3101121333122003): complete subsection reference.

<a id="canonical-3312120121220023-3230131203113023-1222013132133331-2010313033233230-2332333210322111-2102020101101330-2111331013210333-1133201121322223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.snat_pool.no_snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-003.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- [origin_servers.private_name.snat_pool](data-sources--origin_pool--reference--group-003.md#canonical-3312011332131233-1230133131313010-3013003121231323-0223212213120012-0301213303101202-1000312320032330-3210122031212113-2000030222032011)
- origin_servers.private_name.snat_pool.no_snat_pool

<a id="canonical-1231201333130132-1002102100110300-1031113030320232-0200000330233032-0012221031123231-3120101323313200-3122130021321332-3203010113331223"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no snat pool.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3212302120120200-1323200000301103-0223232010213033-3103221203300131-0131233211123212-0100110203310302-1021032313011123-3101121333122003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.private_name.snat_pool.snat_pool` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-003.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- [origin_servers.private_name.snat_pool](data-sources--origin_pool--reference--group-003.md#canonical-3312011332131233-1230133131313010-3013003121231323-0223212213120012-0301213303101202-1000312320032330-3210122031212113-2000030222032011)
- origin_servers.private_name.snat_pool.snat_pool

<a id="canonical-2000120023311112-2013031201210131-3020110013222322-2120332123222321-2232100202201123-3313300131131100-1201032210210303-1012102101332300"></a>

Type: `"single"`. Computed.

List of IPv4 prefixes that represent an endpoint.

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

<a id="canonical-3021110312231313-1013223300320203-2123231113002023-2221132132211320-1013122322203302-2330011123223111-1333032131302103-3301301121030221"></a>

### Direct properties for `origin_servers.private_name.snat_pool.snat_pool`

<a id="canonical-0233103211231133-2110222003232101-0320133301200002-3030012213211320-1033011002033131-3300111023323230-1211310033212312-3120300302301312"></a>

#### `origin_servers.private_name.snat_pool.snat_pool.prefixes` property

Type: `["list", "string"]`. Computed.

List of IPv4 prefixes that represent an endpoint.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 128,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 128,
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
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.items.string.ipv4_prefix": "true",
    "ves.io.schema.rules.repeated.max_items": "128",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3233223332020121-3100021120202320-2031230003101211-3023100123103102-0220333130210323-2201111232211122-1033211313103113-3001033210113321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.public_ip` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- origin_servers.public_ip

<a id="canonical-1123110022211300-0110332003213013-3212002113001201-0213030321132330-0010033001123102-2130022002302130-1032031203320030-0321203102322033"></a>

Type: `"single"`. Computed.

Specify origin server with public IP address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-public_ip_choice": "[\"ip\"]"
}
```

<a id="canonical-2323110231011302-2133120002000113-3310121003110102-2301022313300322-2300033101311230-3021111220132103-3022120223131333-0031031230200000"></a>

### Direct properties for `origin_servers.public_ip`

<a id="canonical-2112313312012323-2310022330233310-3112311113003102-1303230123010002-1331011111322201-0333202020101120-3201102313122030-0100112212232122"></a>

#### `origin_servers.public_ip.ip` property

Type: `"string"`. Computed.

Public IPv4. Exclusive with \[\] Public IPv4 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

<a id="canonical-3023121012300320-1113333033031232-1003130321112112-1203102212132000-1310200323133230-2031320320302202-3121312332032031-2323130012131111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.public_name` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- origin_servers.public_name

<a id="canonical-1233231001013031-2011111132012112-3312030331030032-0110110203203310-3233312102102323-3130123013132101-1330232321023212-2230100130011023"></a>

Type: `"single"`. Computed.

Specify origin server with public DNS name.

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

<a id="canonical-1323210302302033-2302210210110032-0013213302210011-2131220123311320-2213232031123110-2001132112122322-0200231210000300-2231110330123220"></a>

### Direct properties for `origin_servers.public_name`

<a id="canonical-1200212100313122-0230222122110333-3032013211322133-2223123001330111-1103312130303231-1030200121030133-0301323223203322-3302113023330332"></a>

#### `origin_servers.public_name.dns_name` property

Type: `"string"`. Computed.

DNS Name. DNS Name

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 256,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="canonical-3222221303103003-2201330000303113-2100313002211111-0331321033003220-2130311010103001-3200120333100130-3323121003230033-0023233130303021"></a>

<a id="canonical-2012031011312012-2231302022303020-1233021222320211-3332303111222200-0213023101002030-3000002112303301-1300210033230232-2133331211030032"></a>

#### `origin_servers.public_name.refresh_interval` property

Type: `"number"`. Computed.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 604800,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
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
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.ranges": "0,10-604800"
  }
}
```

<a id="canonical-1132312333312222-3332122302332200-2310202220131303-2312323300201122-2030322023323212-2010332131123003-0331231022031330-0313233113021202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.vn_private_ip` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- origin_servers.vn_private_ip

<a id="canonical-1201112113021310-0222213200333021-1033212320200202-3230321213113231-0133112033123230-1022003020032320-2122211222312012-1100330331002123"></a>

Type: `"single"`. Computed.

Specify origin server with IP on Virtual Network.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-virtual_network_ip_choice": "[\"ip\"]"
}
```

<a id="canonical-1213310301123132-0312133022211301-3213021120200000-2001333322300100-1122023321210123-2330311033031310-0102130022132003-3020131033230312"></a>

### Direct properties for `origin_servers.vn_private_ip`

<a id="canonical-2003223332223001-3232311123032310-1312001010002300-2030002132223000-0202011113321011-3220300313123232-2312120221301013-2021130320222030"></a>

#### `origin_servers.vn_private_ip.ip` property

Type: `"string"`. Computed.

IPv4. Exclusive with \[\] IPv4 address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ipv4",
    "maxLength": 1024,
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
    "ves.io.schema.rules.string.ipv4": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ipv4": "true"
  }
}
```

- [virtual_network](data-sources--origin_pool--reference--group-003.md#canonical-2010102112110023-3022011022122033-2323023203111013-3233300030300132-3122331230123120-1200000330211011-0221022301313020-2112231321100023): complete subsection reference.

<a id="canonical-2010102112110023-3022011022122033-2323023203111013-3233300030300132-3122331230123120-1200000330211011-0221022301313020-2112231321100023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.vn_private_ip.virtual_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.vn_private_ip](data-sources--origin_pool--reference--group-003.md#canonical-1132312333312222-3332122302332200-2310202220131303-2312323300201122-2030322023323212-2010332131123003-0331231022031330-0313233113021202)
- origin_servers.vn_private_ip.virtual_network

<a id="canonical-0002021132223321-2003211320303002-1103121321312311-2223210300112331-0300321112200111-3133303122132323-0323330102103021-3030211012303332"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-2010301311333200-1203100322021302-2011300301330010-0033303122003213-1333031222210010-0132031233210300-2100101022122310-2022001230222222"></a>

### Direct properties for `origin_servers.vn_private_ip.virtual_network`

<a id="canonical-1012013330302213-0231110132123330-2133213131023301-2023331220323003-3232332330210031-2130022132022322-2300311020013000-3203103032110201"></a>

#### `origin_servers.vn_private_ip.virtual_network.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3231102213103123-3000012112222310-0203210022120331-2322102213111103-0200311221132311-3320010012311030-1121210020030313-2020301002103123"></a>

<a id="canonical-1221210201032100-3003001223120103-0313232201310310-0122101103103101-0200010130332022-3002022101210100-1310012133311202-1220302211000132"></a>

#### `origin_servers.vn_private_ip.virtual_network.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3312222332103020-2012233203111323-3030232021022110-2312330121102231-1313223332120003-1213102110130332-1022201211112203-3223201102303300"></a>

<a id="canonical-1133031310131112-2000233120221213-0133022200002131-2323112311023233-1201222130211212-2112121310123301-3011200332110120-3230100012010321"></a>

#### `origin_servers.vn_private_ip.virtual_network.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3322112003203211-1300331030110100-3211113003210110-3003332110030121-2200102131333003-3323033022212213-0033103302220322-0021203022320330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.vn_private_name` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- origin_servers.vn_private_name

<a id="canonical-3312300023020300-2221333211123333-2330333201320022-3320130123211313-2112111202221302-2223111321301023-3120123230023022-0031013011113231"></a>

Type: `"single"`. Computed.

Specify origin server with DNS name on Virtual Network.

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

<a id="canonical-0001111201100100-1323321112232332-0212212102333200-3321003121013220-0203113331001132-0020231301201023-3133023013312201-0011021032303132"></a>

### Direct properties for `origin_servers.vn_private_name`

<a id="canonical-1101311331303030-0032012310223201-1030113001021100-3223200320231312-2300221231321110-3211030022002012-0323333220300300-1120312102213313"></a>

#### `origin_servers.vn_private_name.dns_name` property

Type: `"string"`. Computed.

DNS Name. DNS Name

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "network",
    "characterSet": {
      "allowed": "[a-z0-9.-]",
      "description": "Dot-separated DNS labels"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "fqdn",
    "formatDescription": "RFC 1123 FQDN: lowercase, dot-separated labels, max 253 chars total",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.95,
      "source": "inferred",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minLength": 1,
    "pattern": "^([a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?\\.)*[a-z0-9]([a-z0-9-]{0,61}[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1123"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true"
  }
}
```

- [private_network](data-sources--origin_pool--reference--group-003.md#canonical-2022222020213202-3110010323312231-2200233032000223-2213123322022120-1010112202031033-1313212110330102-3021310121230201-1103102210232313): complete subsection reference.

<a id="canonical-2022222020213202-3110010323312231-2200233032000223-2213123322022120-1010112202031033-1313212110330102-3021310121230201-1103102210232313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `origin_servers.vn_private_name.private_network` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-002.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.vn_private_name](data-sources--origin_pool--reference--group-003.md#canonical-3322112003203211-1300331030110100-3211113003210110-3003332110030121-2200102131333003-3323033022212213-0033103302220322-0021203022320330)
- origin_servers.vn_private_name.private_network

<a id="canonical-2022001023201223-1213310013210002-2333312233011020-1133010220111313-0232120012022232-2322002013233023-3321233023222332-3030101333030010"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-2233012333322333-2310301302113132-1303331012001212-2132121323233233-0221230320122021-0101212120032231-3031113200311300-2000313003233023"></a>

### Direct properties for `origin_servers.vn_private_name.private_network`

<a id="canonical-3131030122030032-2223021030311310-3020011103310230-0100010132233300-3332032022303322-0313220320031220-3311102032113020-2333232131200011"></a>

#### `origin_servers.vn_private_name.private_network.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1100122120222203-1201312111313100-1320332211123330-1023231021110333-2322133023220030-2020202003310122-1313211032121220-0133020330002110"></a>

<a id="canonical-2310212220231003-1011323221110330-1032212320033322-0001322301322012-3020122023003200-1120333122231220-0110022330232033-0222133220001210"></a>

#### `origin_servers.vn_private_name.private_network.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-2122133133222333-1331222130230200-3020130132330020-1121001131220300-0230331012033311-3331133333031012-1012302232023131-2121122220201202"></a>

<a id="canonical-0012003213112201-3102001030111132-0000021131203331-2221323113301112-3112000310020202-3331311212211032-0211221331013121-0213232031230013"></a>

#### `origin_servers.vn_private_name.private_network.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3130012011302310-0100000020300011-2333213101123221-2323300231131231-3230213133122123-3221221202130031-0122330100113130-1102320312113033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `same_as_endpoint_port` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- same_as_endpoint_port

<a id="canonical-3230202001231011-0231003131211320-3300230323330013-3310330210031113-0323111321313021-1120103003020320-2233213133003321-0013202003001030"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2231230203233120-1223102033221212-2211002030000333-1332212222223122-0031021313330101-0202130201022222-3131132023113300-0330011232102031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upstream_conn_pool_reuse_type` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- upstream_conn_pool_reuse_type

<a id="canonical-1223310203212231-2312201220222003-0032121130323132-1001322123023120-3030102001303133-0131203102310120-0233101220301003-3210220300301332"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1222110232330200-1312110123003130-0301333113011310-2132122023323032-0211322323023312-2010232203203231-0022023311302022-2230322022101133"></a>

### Direct properties for `upstream_conn_pool_reuse_type`

- [disable_conn_pool_reuse](data-sources--origin_pool--reference--group-003.md#canonical-1022200321320132-1123023332011133-2220201012202003-2200133202320303-0210110310113332-3002003313012203-0131002220211110-2221203011111120): complete subsection reference.

- [enable_conn_pool_reuse](data-sources--origin_pool--reference--group-003.md#canonical-1320101022032011-3203201130311130-0232121211022033-3203300130321003-3323120121012121-0112332312130333-0000133030030300-2012001123130123): complete subsection reference.

<a id="canonical-1022200321320132-1123023332011133-2220201012202003-2200133202320303-0210110310113332-3002003313012203-0131002220211110-2221203011111120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upstream_conn_pool_reuse_type.disable_conn_pool_reuse` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [upstream_conn_pool_reuse_type](data-sources--origin_pool--reference--group-003.md#canonical-2231230203233120-1223102033221212-2211002030000333-1332212222223122-0031021313330101-0202130201022222-3131132023113300-0330011232102031)
- upstream_conn_pool_reuse_type.disable_conn_pool_reuse

<a id="canonical-1221113302121301-2313020121331320-0221203333123323-2123200223300002-1330031110322133-3313200133100312-1310111322321002-2123221332010213"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1320101022032011-3203201130311130-0232121211022033-3203300130321003-3323120121012121-0112332312130333-0000133030030300-2012001123130123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `upstream_conn_pool_reuse_type.enable_conn_pool_reuse` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [upstream_conn_pool_reuse_type](data-sources--origin_pool--reference--group-003.md#canonical-2231230203233120-1223102033221212-2211002030000333-1332212222223122-0031021313330101-0202130201022222-3131132023113300-0330011232102031)
- upstream_conn_pool_reuse_type.enable_conn_pool_reuse

<a id="canonical-1303312032131320-2021212201121011-2021212032121302-0213111331132120-0201200332100220-0201322223303113-2130120213131233-3303120103001132"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- use_tls

<a id="canonical-0002222103101132-0331331011030000-0221212201230213-1311030121013132-0223020023313332-2311112031102200-3132032023331103-2122132001320021"></a>

Type: `"single"`. Computed.

TLS Parameters for Origin Servers. Upstream TLS Parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-max_session_keys_type": "[\"default_session_key_caching\",\"disable_session_key_caching\",\"max_session_keys\"]",
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\",\"use_mtls_obj\"]",
  "x-ves-oneof-field-server_validation_choice": "[\"skip_server_verification\",\"use_server_verification\",\"volterra_trusted_ca\"]",
  "x-ves-oneof-field-sni_choice": "[\"disable_sni\",\"sni\",\"use_host_header_as_sni\"]"
}
```

<a id="canonical-1220221332111002-3221322023100232-3112031001000133-0231121231332321-1202012120331222-0013133323002211-2131203222110102-2231231000003222"></a>

### Direct properties for `use_tls`

- [default_session_key_caching](data-sources--origin_pool--reference--group-003.md#canonical-3120033223300122-0002012021322010-3133211003121300-2201320200112013-1022121320122111-1321231300011022-1202030223010131-2130002001033012): complete subsection reference.

- [disable_session_key_caching](data-sources--origin_pool--reference--group-003.md#canonical-3112023032132312-1132010002311002-3003320230110132-1123303202301332-0312123102222332-0303130310001332-1123133313201330-2001030020320211): complete subsection reference.

- [disable_sni](data-sources--origin_pool--reference--group-003.md#canonical-3302331001011103-0032210313332130-2330101100323023-3233031132332300-3233222030222021-0231223303003021-1220133302031312-1211313003212132): complete subsection reference.

<a id="canonical-0233103101220320-1020311303110113-1100311111211321-1231332102301202-2320032130222232-0033111113112220-2100121012331321-0130112230133010"></a>

<a id="canonical-0013302310100031-3211121021101022-0100131102103311-2321131313002001-1102232200012121-0233301101233130-3323210312130232-2213303011030310"></a>

#### `use_tls.max_session_keys` property

Type: `"number"`. Computed.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 64,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minimum": 2
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "2",
    "ves.io.schema.rules.uint32.lte": "64"
  }
}
```

- [no_mtls](data-sources--origin_pool--reference--group-003.md#canonical-2033030130120230-2002013012023100-2300101011010123-0023110121222313-2002022101100132-1200212102011022-0232110202003010-1021021323033132): complete subsection reference.

- [skip_server_verification](data-sources--origin_pool--reference--group-003.md#canonical-3133102211002303-3032200122321120-0210020232102320-2021111001013011-1213213100231100-2122021233030322-3321111131233310-0313010303321021): complete subsection reference.

<a id="canonical-3313012030230223-0020202211000122-2020100322223211-2300020233303232-3202222230130322-1233213120013233-0002012022211210-0321132221022123"></a>

<a id="canonical-0123001130223101-0321300122302323-2200010121121213-1131021231333302-0020131300000322-3331112312122331-2100223301303113-3312310131312200"></a>

#### `use_tls.sni` property

Type: `"string"`. Computed.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 256,
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
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "256"
  }
}
```

- [tls_config](data-sources--origin_pool--reference--group-003.md#canonical-3002221101212230-2032212202012222-2211102302333323-3031111211101332-0202202331023033-3323013021210110-1130132200330201-1122230100213213): complete subsection reference.

- [use_host_header_as_sni](data-sources--origin_pool--reference--group-003.md#canonical-3232122223122133-3321210122301001-0222031200332131-2203033233132002-0222102221311230-2103030133101313-3220313220123111-1321233321311021): complete subsection reference.

- [use_mtls](data-sources--origin_pool--reference--group-003.md#canonical-2123230122313111-0311121021121133-1321002130011032-3312300200211002-2110033332211203-2121202220332232-0100223311202100-1321101100003310): complete subsection reference.

- [use_mtls_obj](data-sources--origin_pool--reference--group-003.md#canonical-3310130231121301-1123130000302011-2213030223021110-2221211022020003-3033112020011311-2022011333233203-1120313312223010-3003101030232320): complete subsection reference.

- [use_server_verification](data-sources--origin_pool--reference--group-003.md#canonical-3013231000221313-0112303320203120-0231130102303320-1100002233231301-1222203202030202-1113323222002130-0331312102131100-0102330131222110): complete subsection reference.

- [volterra_trusted_ca](data-sources--origin_pool--reference--group-003.md#canonical-1332301321201201-2112102322122113-3211233001322320-1230113120130000-0130212221102210-3323230123233332-1121303303011030-0131110111202333): complete subsection reference.

<a id="canonical-3120033223300122-0002012021322010-3133211003121300-2201320200112013-1022121320122111-1321231300011022-1202030223010131-2130002001033012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.default_session_key_caching` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- use_tls.default_session_key_caching

<a id="canonical-2322202020222130-1232110023322133-3333101113001113-3210232200231331-1231120123332320-3010301033212312-3103012211310233-0311213210031323"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default session key caching. Defaults to \`map\[\]\`. Server applies
default when omitted.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3112023032132312-1132010002311002-3003320230110132-1123303202301332-0312123102222332-0303130310001332-1123133313201330-2001030020320211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.disable_session_key_caching` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- use_tls.disable_session_key_caching

<a id="canonical-2003010311021311-0332020301111223-3123232101102112-3103113001010103-0333200233003302-2102011021303311-1112312022322313-1302011130230300"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3302331001011103-0032210313332130-2330101100323023-3233031132332300-3233222030222021-0231223303003021-1220133302031312-1211313003212132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.disable_sni` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- use_tls.disable_sni

<a id="canonical-1212122200212010-3013213213022221-3021312132133231-1021313013010111-3031231222111033-0211300322212000-2333213222123033-2220012003130200"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2033030130120230-2002013012023100-2300101011010123-0023110121222313-2002022101100132-1200212102011022-0232110202003010-1021021323033132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.no_mtls` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- use_tls.no_mtls

<a id="canonical-2121302331023120-3120231203230320-1111220312012332-1030122211200213-2330101321313131-1023131222221102-0300313213311031-0111103120033123"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3133102211002303-3032200122321120-0210020232102320-2021111001013011-1213213100231100-2122021233030322-3321111131233310-0313010303321021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.skip_server_verification` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- use_tls.skip_server_verification

<a id="canonical-3331020311210023-1121220230320302-1302202211103233-2310322312021232-3132220231332313-3321331123311103-3200303312312113-2001130112031321"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3002221101212230-2032212202012222-2211102302333323-3031111211101332-0202202331023033-3323013021210110-1130132200330201-1122230100213213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.tls_config` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- use_tls.tls_config

<a id="canonical-2221121021212003-3311211133031031-3322231011231023-2021100112101121-1130032323313321-1022313211131100-1033020031230122-2332321020033203"></a>

Type: `"single"`. Computed.

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "tls",
    "constraintType": "object",
    "deterministic": true,
    "metadata": {
      "category": "tls",
      "confidence": 0.99,
      "note": "Required when use_tls is selected — API returns 400 if nil",
      "source": "api-probed",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

<a id="canonical-1202332303021122-0011211232000002-0122201310011200-3000011213310031-0332230321002301-1013210222010100-2322332112103031-3022330331331332"></a>

### Direct properties for `use_tls.tls_config`

- [custom_security](data-sources--origin_pool--reference--group-003.md#canonical-1330233100130232-3232102231013300-0032311102013312-1222000110023213-0200330121002332-3302110021020222-0033012322220102-0213123123130201): complete subsection reference.

- [default_security](data-sources--origin_pool--reference--group-003.md#canonical-3201003131323002-3001122131100101-1213030220013230-3100320313122031-3201013213203130-2002300330233211-3133132022320332-0001232012203310): complete subsection reference.

- [low_security](data-sources--origin_pool--reference--group-003.md#canonical-2122300232331021-3313012021313120-0202322113220311-0000000021223221-3002333120022322-3321030021122020-2302122013233112-1130211002120010): complete subsection reference.

- [medium_security](data-sources--origin_pool--reference--group-003.md#canonical-3202310210131201-2313133032112333-0020132122322123-1101111211101001-2231200021132130-0003200102332022-3332333222033101-3122121200213221): complete subsection reference.

<a id="canonical-1330233100130232-3232102231013300-0032311102013312-1222000110023213-0200330121002332-3302110021020222-0033012322220102-0213123123130201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.tls_config.custom_security` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- [use_tls.tls_config](data-sources--origin_pool--reference--group-003.md#canonical-3002221101212230-2032212202012222-2211102302333323-3031111211101332-0202202331023033-3323013021210110-1130132200330201-1122230100213213)
- use_tls.tls_config.custom_security

<a id="canonical-3103212201032200-3132123333012101-3303303013300030-2113221030022312-3111111231031111-1130120113103330-2101111111302113-3003322113131103"></a>

Type: `"single"`. Computed.

This defines TLS protocol config including min/max versions and allowed ciphers.

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

<a id="canonical-3001033323333132-3120031013113310-1131100003320211-2233123022203331-0121200121033222-3223313211223100-1003300111113033-2310021230033320"></a>

### Direct properties for `use_tls.tls_config.custom_security`

<a id="canonical-1133310020130221-0011030031032112-0230300021102011-3333000223122333-2102211003312202-3033132231022221-3013102330002300-3113023000223230"></a>

#### `use_tls.tls_config.custom_security.cipher_suites` property

Type: `["list", "string"]`. Computed.

The TLS listener will only support the specified cipher list.

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
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.items.string.in": "[\\\"TLS_AES_128_GCM_SHA256\\\",\\\"TLS_AES_256_GCM_SHA384\\\",\\\"TLS_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384\\\",\\\"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_128_GCM_SHA256\\\",\\\"TLS_RSA_WITH_AES_256_CBC_SHA\\\",\\\"TLS_RSA_WITH_AES_256_GCM_SHA384\\\"]",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

<a id="canonical-3323313112120033-2112103233111132-3132310012221221-0100020120032033-1221231132003020-3012032210330030-0023120100120310-3223220112113200"></a>

<a id="canonical-1113211320321322-0321222330133031-1020113100103322-2000002123232210-2303111031100311-3201310113001101-2201022130010213-3123332330111123"></a>

#### `use_tls.tls_config.custom_security.max_version` property

Type: `"string"`. Computed.

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

<a id="canonical-1222322212103203-3310011001332312-2233323112233022-0232112103330322-3221101221022300-0132231221031332-3032330132002131-0012002231200003"></a>

<a id="canonical-2011200232011212-3110231033023130-0201232232201012-0010300321332233-1322120210101002-3023231300331232-3231132110000012-0013203132211333"></a>

#### `use_tls.tls_config.custom_security.min_version` property

Type: `"string"`. Computed.

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

<a id="canonical-3201003131323002-3001122131100101-1213030220013230-3100320313122031-3201013213203130-2002300330233211-3133132022320332-0001232012203310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.tls_config.default_security` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- [use_tls.tls_config](data-sources--origin_pool--reference--group-003.md#canonical-3002221101212230-2032212202012222-2211102302333323-3031111211101332-0202202331023033-3323013021210110-1130132200330201-1122230100213213)
- use_tls.tls_config.default_security

<a id="canonical-3011131100330232-2331223001030001-2100020233001001-0033112210302331-3123310210321131-3023221333003230-3102230110302311-3332321000321101"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2122300232331021-3313012021313120-0202322113220311-0000000021223221-3002333120022322-3321030021122020-2302122013233112-1130211002120010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.tls_config.low_security` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- [use_tls.tls_config](data-sources--origin_pool--reference--group-003.md#canonical-3002221101212230-2032212202012222-2211102302333323-3031111211101332-0202202331023033-3323013021210110-1130132200330201-1122230100213213)
- use_tls.tls_config.low_security

<a id="canonical-3001030021000000-0123220133112320-3002232333203323-2120131120102120-3230310331232212-1211131130102033-3302113330313330-2202310312031031"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3202310210131201-2313133032112333-0020132122322123-1101111211101001-2231200021132130-0003200102332022-3332333222033101-3122121200213221"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.tls_config.medium_security` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- [use_tls.tls_config](data-sources--origin_pool--reference--group-003.md#canonical-3002221101212230-2032212202012222-2211102302333323-3031111211101332-0202202331023033-3323013021210110-1130132200330201-1122230100213213)
- use_tls.tls_config.medium_security

<a id="canonical-1231213011331323-3030123110201313-3103022200330311-3312021113003211-2123033032331133-2323011323210000-1103310323202321-1132012001303232"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3232122223122133-3321210122301001-0222031200332131-2203033233132002-0222102221311230-2103030133101313-3220313220123111-1321233321311021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_host_header_as_sni` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- use_tls.use_host_header_as_sni

<a id="canonical-3002002300202330-0332220022012011-3022022020123221-1033202333233112-3210233212112023-3210131231311110-1212330303233012-0223130102000013"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2123230122313111-0311121021121133-1321002130011032-3312300200211002-2110033332211203-2121202220332232-0100223311202100-1321101100003310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- use_tls.use_mtls

<a id="canonical-0033031012120110-2120022223112020-3023203121021012-3023302113320212-1232031232003020-0312120013033313-1003003313101323-1112303001113332"></a>

Type: `"single"`. Computed.

MTLS Certificate. MTLS Client Certificate.

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

<a id="canonical-2222303221132123-0003033222312220-0212001102103132-1132313331122200-1112313013100201-2112333023003300-2200012111320300-1220233222111110"></a>

### Direct properties for `use_tls.use_mtls`

- [tls_certificates](data-sources--origin_pool--reference--group-003.md#canonical-1303112212030232-0223201010130110-2203010323313333-1333323310300102-0010003103112021-3212312311323231-0313030221021313-3302023303112222): complete subsection reference.

<a id="canonical-1303112212030232-0223201010130110-2203010323313333-1333323310300102-0010003103112021-3212312311323231-0313030221021313-3302023303112222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls.tls_certificates` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- [use_tls.use_mtls](data-sources--origin_pool--reference--group-003.md#canonical-2123230122313111-0311121021121133-1321002130011032-3312300200211002-2110033332211203-2121202220332232-0100223311202100-1321101100003310)
- use_tls.use_mtls.tls_certificates

<a id="canonical-2202212233123011-2303103030233031-3000003203010000-2300010200213231-3013301111123220-0312101101002122-1230132311310123-1321203332102332"></a>

Type: `"list"`. Computed.

MTLS Client Certificate. MTLS Client Certificate.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 1,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 1,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-10T03:47:03+00:00"
    },
    "minItems": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "1",
    "ves.io.schema.rules.repeated.min_items": "1"
  }
}
```

<a id="canonical-1213203320332323-2131200110123111-0100113230311010-0333310023121203-0231201113031011-0020021311111032-3231122122231210-3112012302213012"></a>

### Direct properties for `use_tls.use_mtls.tls_certificates`

<a id="canonical-1223331010111320-0112301120111003-3312112000101332-3233130011303312-1211233011232113-2130100220213221-2113133230300321-2310100030001012"></a>

#### `use_tls.use_mtls.tls_certificates.certificate_url` property

Type: `"string"`. Computed.

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

- [custom_hash_algorithms](data-sources--origin_pool--reference--group-003.md#canonical-2213022133103033-2013320233231103-3020122200330020-2312330312113331-0221000331300331-1210220320322122-2222002203221332-0213302200012202): complete subsection reference.

<a id="canonical-1031213332303213-1220020112222332-2300302013112023-3012013211023210-1023331301303220-3232002013123311-3121321102320012-3322122223210321"></a>

<a id="canonical-3223312311011133-3123300222210100-0003210330223220-0212131123022022-1010113323013113-1022120213121101-0132023203013311-2013100301223021"></a>

#### `use_tls.use_mtls.tls_certificates.description_spec` property

Type: `"string"`. Computed.

Description. Description for the certificate.

- [disable_ocsp_stapling](data-sources--origin_pool--reference--group-003.md#canonical-2120212310332011-2032230310332212-0320121121231023-1222202302203313-3302311300301111-2130202202130220-1120101110022211-3110200232030130): complete subsection reference.

- [private_key](data-sources--origin_pool--reference--group-003.md#canonical-1302022202330221-3120232102301233-0210130301320303-2200223232022212-1303201200031330-3210030221113213-0311311312222333-0322101101323212): complete subsection reference.

- [use_system_defaults](data-sources--origin_pool--reference--group-003.md#canonical-1013131312123031-0232130211120013-3322110220233232-2003322332003002-0121213122132302-2022211132100130-0232303021231030-3332230313331211): complete subsection reference.

<a id="canonical-2213022133103033-2013320233231103-3020122200330020-2312330312113331-0221000331300331-1210220320322122-2222002203221332-0213302200012202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls.tls_certificates.custom_hash_algorithms` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- [use_tls.use_mtls](data-sources--origin_pool--reference--group-003.md#canonical-2123230122313111-0311121021121133-1321002130011032-3312300200211002-2110033332211203-2121202220332232-0100223311202100-1321101100003310)
- [use_tls.use_mtls.tls_certificates](data-sources--origin_pool--reference--group-003.md#canonical-1303112212030232-0223201010130110-2203010323313333-1333323310300102-0010003103112021-3212312311323231-0313030221021313-3302023303112222)
- use_tls.use_mtls.tls_certificates.custom_hash_algorithms

<a id="canonical-1220102320202303-1111031101221022-0202303032210332-0131203011230003-0113113012202331-0223020030013103-3123203223302003-2320000301233032"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3222123012200212-1220200113021001-3010320331302030-3321311103302002-3030010212123203-2201233213302100-2021220001002231-3321220312111113"></a>

### Direct properties for `use_tls.use_mtls.tls_certificates.custom_hash_algorithms`

<a id="canonical-1220012221230213-1013230221122030-0112110300201302-3132001013300213-0232133123213203-1230022011111103-0323120103032330-3033101021002000"></a>

#### `use_tls.use_mtls.tls_certificates.custom_hash_algorithms.hash_algorithms` property

Type: `["list", "string"]`. Computed.

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

<a id="canonical-2120212310332011-2032230310332212-0320121121231023-1222202302203313-3302311300301111-2130202202130220-1120101110022211-3110200232030130"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls.tls_certificates.disable_ocsp_stapling` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- [use_tls.use_mtls](data-sources--origin_pool--reference--group-003.md#canonical-2123230122313111-0311121021121133-1321002130011032-3312300200211002-2110033332211203-2121202220332232-0100223311202100-1321101100003310)
- [use_tls.use_mtls.tls_certificates](data-sources--origin_pool--reference--group-003.md#canonical-1303112212030232-0223201010130110-2203010323313333-1333323310300102-0010003103112021-3212312311323231-0313030221021313-3302023303112222)
- use_tls.use_mtls.tls_certificates.disable_ocsp_stapling

<a id="canonical-0110213323332312-3031023021013130-2122032320303101-3032133230023111-3223111320302313-1332022212313312-3230033031310223-1300123101012211"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1302022202330221-3120232102301233-0210130301320303-2200223232022212-1303201200031330-3210030221113213-0311311312222333-0322101101323212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls.tls_certificates.private_key` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- [use_tls.use_mtls](data-sources--origin_pool--reference--group-003.md#canonical-2123230122313111-0311121021121133-1321002130011032-3312300200211002-2110033332211203-2121202220332232-0100223311202100-1321101100003310)
- [use_tls.use_mtls.tls_certificates](data-sources--origin_pool--reference--group-003.md#canonical-1303112212030232-0223201010130110-2203010323313333-1333323310300102-0010003103112021-3212312311323231-0313030221021313-3302023303112222)
- use_tls.use_mtls.tls_certificates.private_key

<a id="canonical-2301130302322321-3131122203022311-3103012320011131-1202210201223201-2202202203122120-0030102031112032-0332003000032301-2012103133130013"></a>

Type: `"single"`. Computed.

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

<a id="canonical-3003113133030300-2132120323312202-0222233122113032-1031033313200031-0031002321030121-0322202132013231-2111023310012021-0320301031003120"></a>

### Direct properties for `use_tls.use_mtls.tls_certificates.private_key`

- [blindfold_secret_info](data-sources--origin_pool--reference--group-003.md#canonical-1321101000102310-2213300110023332-2231103110223303-2100103301222323-0123201223000113-1210200332301321-1031000221130302-0002110323000101): complete subsection reference.

- [clear_secret_info](data-sources--origin_pool--reference--group-003.md#canonical-0103020332012300-2232300202231330-1113110320021021-1133213310010003-3001303110300131-3311221213002011-0232232212020123-2123301112220210): complete subsection reference.

<a id="canonical-1321101000102310-2213300110023332-2231103110223303-2100103301222323-0123201223000113-1210200332301321-1031000221130302-0002110323000101"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- [use_tls.use_mtls](data-sources--origin_pool--reference--group-003.md#canonical-2123230122313111-0311121021121133-1321002130011032-3312300200211002-2110033332211203-2121202220332232-0100223311202100-1321101100003310)
- [use_tls.use_mtls.tls_certificates](data-sources--origin_pool--reference--group-003.md#canonical-1303112212030232-0223201010130110-2203010323313333-1333323310300102-0010003103112021-3212312311323231-0313030221021313-3302023303112222)
- [use_tls.use_mtls.tls_certificates.private_key](data-sources--origin_pool--reference--group-003.md#canonical-1302022202330221-3120232102301233-0210130301320303-2200223232022212-1303201200031330-3210030221113213-0311311312222333-0322101101323212)
- use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info

<a id="canonical-0023300230121122-1013012102031220-0033103301023100-0321233132332203-3332230001301100-2101012100113110-0333113013221301-0222223230111130"></a>

Type: `"single"`. Computed.

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

<a id="canonical-2130221102320311-3333112231112202-3300030020302033-0223330221113302-1030032221012332-3303101120012100-3101013001102310-2323320300300210"></a>

### Direct properties for `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info`

<a id="canonical-0301303131210231-1022000223200132-0001121001100210-3222333213221002-1330330323300021-0331311003102212-2321223003220232-3321221223012130"></a>

#### `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.decryption_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-1000001120232231-2000101030032122-2203213222130300-1230232222111222-0221012231233210-3113302301013310-2311302213223120-2023102220021333"></a>

<a id="canonical-3003013012213121-3311312032232332-2103122122320201-3230320100212122-2301022230220023-3102222212012031-2232232203301122-2201330100300230"></a>

#### `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.location` property

Type: `"string"`. Computed, Sensitive.

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

<a id="canonical-2300023330023010-0210211132203102-3013321003332021-2323201003212200-1130200321302102-1201030031313221-2112110300103320-0202203123212301"></a>

<a id="canonical-0211321313030012-3323303301320031-1023131302011200-1302302200220221-0223020103310303-1000102112011212-0200200222122002-2103311102222312"></a>

#### `use_tls.use_mtls.tls_certificates.private_key.blindfold_secret_info.store_provider` property

Type: `"string"`. Computed.

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

<a id="canonical-0103020332012300-2232300202231330-1113110320021021-1133213310010003-3001303110300131-3311221213002011-0232232212020123-2123301112220210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- [use_tls.use_mtls](data-sources--origin_pool--reference--group-003.md#canonical-2123230122313111-0311121021121133-1321002130011032-3312300200211002-2110033332211203-2121202220332232-0100223311202100-1321101100003310)
- [use_tls.use_mtls.tls_certificates](data-sources--origin_pool--reference--group-003.md#canonical-1303112212030232-0223201010130110-2203010323313333-1333323310300102-0010003103112021-3212312311323231-0313030221021313-3302023303112222)
- [use_tls.use_mtls.tls_certificates.private_key](data-sources--origin_pool--reference--group-003.md#canonical-1302022202330221-3120232102301233-0210130301320303-2200223232022212-1303201200031330-3210030221113213-0311311312222333-0322101101323212)
- use_tls.use_mtls.tls_certificates.private_key.clear_secret_info

<a id="canonical-2303113123201301-2030130200000221-3012311131022123-2330101110113102-2023122210031212-3100000212000200-1302232322201302-3131210211211130"></a>

Type: `"single"`. Computed.

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

<a id="canonical-1300200021302010-2323201001310132-2120001330202022-1232210330022001-0230203013130032-3102002112010310-3032201203312210-1002301200111023"></a>

### Direct properties for `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info`

<a id="canonical-0110120331231132-2020230322002312-0300020030303202-0201330303232203-3233221332133102-2232212310200111-1230221212331322-2123032321031122"></a>

#### `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.provider_ref` property

Type: `"string"`. Computed.

Name of the Secret Management Access object that contains information about the store to GET
encrypted bytes This field needs to be provided only if the URL scheme is not string:///.

<a id="canonical-0232203120222331-2123332013313223-1031102202300012-2101011031220220-0213203110223213-3033322203111102-0122103332112032-1331201233302323"></a>

<a id="canonical-2230323002312121-2103212011213203-2300200201311211-2321233022022022-3202313111202312-1133330232020120-0302122100111202-2000231233013333"></a>

#### `use_tls.use_mtls.tls_certificates.private_key.clear_secret_info.url` property

Type: `"string"`. Computed, Sensitive.

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

<a id="canonical-1013131312123031-0232130211120013-3322110220233232-2003322332003002-0121213122132302-2022211132100130-0232303021231030-3332230313331211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls.tls_certificates.use_system_defaults` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- [use_tls.use_mtls](data-sources--origin_pool--reference--group-003.md#canonical-2123230122313111-0311121021121133-1321002130011032-3312300200211002-2110033332211203-2121202220332232-0100223311202100-1321101100003310)
- [use_tls.use_mtls.tls_certificates](data-sources--origin_pool--reference--group-003.md#canonical-1303112212030232-0223201010130110-2203010323313333-1333323310300102-0010003103112021-3212312311323231-0313030221021313-3302023303112222)
- use_tls.use_mtls.tls_certificates.use_system_defaults

<a id="canonical-0031310212103130-1310313100023112-3012131302223101-1321300211310021-2300130120123321-1010223012202000-1113010232212021-3231031110032022"></a>

Type: `["object", {}]`. Computed.

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

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3310130231121301-1123130000302011-2213030223021110-2221211022020003-3033112020011311-2022011333233203-1120313312223010-3003101030232320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_mtls_obj` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- use_tls.use_mtls_obj

<a id="canonical-3221232002222233-1031100031012131-0031311201022132-1102211130230203-1110312313123220-2203333133012202-3120212111302220-0210312301230002"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-1013110210113311-0223200013132012-1000120120121013-0120022300031322-2030213320230332-2220313032210332-0301223002100301-3201110133022330"></a>

### Direct properties for `use_tls.use_mtls_obj`

<a id="canonical-0131030113033223-1123123110001032-1333303333212002-2131322113223303-0201312222332331-0222001100132022-0103320010131001-2022021003103010"></a>

#### `use_tls.use_mtls_obj.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-1100220001210103-1322301223000031-0030301333010321-0000201011023233-2122131202231120-2010120001133301-3232233000220220-0122212310313223"></a>

<a id="canonical-2112111112223021-1221310312310021-2201231111322002-2022222112001102-2130130022222220-3003302031113223-2011213133112122-1100112131113323"></a>

#### `use_tls.use_mtls_obj.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0300310203232220-0033212011213220-3023322010122113-3231202031203032-2320330230102210-1031213123133131-2032100002222201-3013123300233213"></a>

<a id="canonical-2200011003022001-0023332012101110-0223132330031232-3110232202333333-0311232313311002-3030312233120002-0302112132233321-0131331010330112"></a>

#### `use_tls.use_mtls_obj.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-3013231000221313-0112303320203120-0231130102303320-1100002233231301-1222203202030202-1113323222002130-0331312102131100-0102330131222110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_server_verification` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- use_tls.use_server_verification

<a id="canonical-2332232310221103-3020011210223113-3010300002012100-2220313031000030-0330012230320200-3300333103111231-3201321020123012-3101212223012013"></a>

Type: `"single"`. Computed.

Configuration parameter for use server verification.

Additional upstream details:

Upstream TLS Validation Context.

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

<a id="canonical-3221232332020220-3201031203131033-1200311110312101-2002311312030200-3332301022330203-0312102332210213-0013111023021133-3321313100333113"></a>

### Direct properties for `use_tls.use_server_verification`

- [trusted_ca](data-sources--origin_pool--reference--group-003.md#canonical-1001012100031011-3213231130020211-1320002102223002-1031230111320030-1100210002212203-3302203022111120-0012321313330002-0123322221310102): complete subsection reference.

<a id="canonical-2112300001303223-1333103223233110-2022201121123023-3131010003320103-1112123121311320-0312231133021030-3010010331310213-0000122010133112"></a>

<a id="canonical-2130110011320311-1212201121312303-3001333200130130-1111131233132011-2231122220223322-3002022102300230-0320011003032332-2031032210000332"></a>

#### `use_tls.use_server_verification.trusted_ca_url` property

Type: `"string"`. Computed.

Exclusive with \[trusted\_ca\] Upload a Root CA Certificate specifically for this Origin Pool for
verification of server's certificate.

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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "131072",
    "ves.io.schema.rules.string.min_bytes": "1",
    "ves.io.schema.rules.string.truststore_url": "true"
  }
}
```

<a id="canonical-1001012100031011-3213231130020211-1320002102223002-1031230111320030-1100210002212203-3302203022111120-0012321313330002-0123322221310102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.use_server_verification.trusted_ca` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- [use_tls.use_server_verification](data-sources--origin_pool--reference--group-003.md#canonical-3013231000221313-0112303320203120-0231130102303320-1100002233231301-1222203202030202-1113323222002130-0331312102131100-0102330131222110)
- use_tls.use_server_verification.trusted_ca

<a id="canonical-2333330121103121-1211202222332020-2220102122231311-0133230102301211-0202001303010211-0101212000031101-2330211321000301-1311330112030110"></a>

Type: `"single"`. Computed.

This type establishes a direct reference from one object(the referrer) to another(the referred).
Such a reference is in form of tenant/namespace/name.

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

<a id="canonical-2101133321201121-0111112303202302-2033201323012033-2030320123231210-0001211003103120-1311122131223102-0123103010310200-1101300012212232"></a>

### Direct properties for `use_tls.use_server_verification.trusted_ca`

<a id="canonical-3101010321130322-1101022122111133-1122313113112132-3131001013313230-2300230300220313-0123203031320000-3311132120012022-3120030313020001"></a>

#### `use_tls.use_server_verification.trusted_ca.name` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 1,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 128,
      "min": 1
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
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
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_bytes": "128",
    "ves.io.schema.rules.string.min_bytes": "1"
  }
}
```

<a id="canonical-3023211323021121-3302133322231222-0013230011320232-2333121231111320-0323312020212331-3311302300212133-3121003022101030-3222100131032332"></a>

<a id="canonical-1030221111012003-0233020020223230-3213011222233131-1303032300203100-3300031230023210-2320203023202011-3103003300301032-0021330102230130"></a>

#### `use_tls.use_server_verification.trusted_ca.namespace` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
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
      "source": "discovery",
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
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-0131102013021100-3200112220313010-2112013330001312-0100201302000030-2332220112203202-1101331021302131-2333000222130200-0320213311132131"></a>

<a id="canonical-1000001010111330-1102332231303021-0003012023332012-1223002212110133-0020112331232102-2003031302312000-1112212302231022-2313032331210312"></a>

#### `use_tls.use_server_verification.trusted_ca.tenant` property

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 64,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 64
    },
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 64,
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
    "ves.io.schema.rules.string.max_bytes": "64"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.max_bytes": "64"
  }
}
```

<a id="canonical-1332301321201201-2112102322122113-3211233001322320-1230113120130000-0130212221102210-3323230123233332-1121303303011030-0131110111202333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## `use_tls.volterra_trusted_ca` properties

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-003.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- use_tls.volterra_trusted_ca

<a id="canonical-2312213023123311-1023103300000102-3313110223220323-1212011230113131-3201101233232303-3110122023021312-1010000301122132-2011300003022123"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for volterra trusted ca. Defaults to \`map\[\]\`. Server applies default
when omitted.

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

This is an empty object or choice marker. It has no direct properties.
