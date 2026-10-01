---
page_title: "xcsh_origin_pool reference"
subcategory: "Load Balancing"
description: "Complete grouped canonical reference for xcsh_origin_pool reference."
---

# xcsh_origin_pool reference

<a id="canonical-1102021123131220-3303000221322011-0233210020030231-3013003221001022-1003013013311120-2001200312022202-1110233330201303-0203131311131310"></a>

## Direct properties — consul_service / 222303003123 / 3

- [inside_network](data-sources--origin_pool--reference--group-002.md#canonical-1311221233101121-0021312202000313-3231111332102020-0021232221101013-3101223120103103-3130111102132330-3132232230311213-0110113113121122): complete subsection reference.

- [outside_network](data-sources--origin_pool--reference--group-002.md#canonical-1130022023121110-3213212022312101-3333221030031123-1321022032300221-0231103100232131-0200212110210030-1202311120300013-1232312120303113): complete subsection reference.

<a id="canonical-1300121122331122-1123001012133210-2000202031321320-1323001021210011-1211001310113113-1230021010000203-2312313322320211-1112211122322231"></a>

<a id="canonical-1002112021030210-1113231313132100-2113232311033102-3213322222203212-0110311232100011-0310100121121030-1313003130022123-1021320010302303"></a>

## service_name property — consul_service / 222303003123 / 4

Type: `"string"`. Computed.

Consul service name of this origin server will be listed, including cluster-ID. The format is
servicename:cluster-ID.

Upstream description:

Consul service name of this origin server will be listed, including cluster-ID. The format is
servicename:cluster-ID.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [site_locator](data-sources--origin_pool--reference--group-002.md#canonical-0220232232031101-0021332122202322-0000010021303022-3120030132133221-3332012010321202-3213003213331110-3022311110231022-0030230300222112): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-2320132102213312-1002323332032110-2303031210111010-2303200111311303-3103313332001302-2113122313133300-0233103020312222-0232102301003021): complete subsection reference.

<a id="canonical-3223120121103323-1200100311031203-2210103322112011-2301212112320002-2033221203223122-3312331011011230-0001313112222231-1011100123113101"></a>

## Next pages — consul_service / 222303003123 / 5

- [origin_servers.consul_service.inside_network](data-sources--origin_pool--reference--group-002.md#canonical-1311221233101121-0021312202000313-3231111332102020-0021232221101013-3101223120103103-3130111102132330-3132232230311213-0110113113121122)
- [origin_servers.consul_service.outside_network](data-sources--origin_pool--reference--group-002.md#canonical-1130022023121110-3213212022312101-3333221030031123-1321022032300221-0231103100232131-0200212110210030-1202311120300013-1232312120303113)
- [origin_servers.consul_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-0220232232031101-0021332122202322-0000010021303022-3120030132133221-3332012010321202-3213003213331110-3022311110231022-0030230300222112)
- [origin_servers.consul_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-2320132102213312-1002323332032110-2303031210111010-2303200111311303-3103313332001302-2113122313133300-0233103020312222-0232102301003021)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-1311221233101121-0021312202000313-3231111332102020-0021232221101013-3101223120103103-3130111102132330-3132232230311213-0110113113121122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2131031323201012-1033013203223303-0010103031301323-0113313300312130-0101331322012211-2312313113021323-1013323102132121-2313310231120021"></a>

## origin_servers.consul_service.inside_network — inside_network / 331211203023 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101)
- origin_servers.consul_service.inside_network

<a id="canonical-1320231211131113-0312330100313120-0000133233102123-0232312210131300-1202121323203102-3332212312233131-0332223332023321-0113223122313111"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside network.

Upstream description:

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

<a id="canonical-2301301111022323-0233023100322220-3231021110100310-2201221333121123-0002002113332023-3203322303103112-3112201020112110-3223133223123031"></a>

## Direct properties — inside_network / 331211203023 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0021332230212200-1021011120130123-0321203003222302-1301133023031030-0103022313313302-0220023102100133-3212133313000303-0210112313022330"></a>

## Next pages — inside_network / 331211203023 / 4

- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-1130022023121110-3213212022312101-3333221030031123-1321022032300221-0231103100232131-0200212110210030-1202311120300013-1232312120303113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1232213301223330-0120123020032032-2101200303221121-2200133323331111-3031021302201030-3213130130211202-0313032131131332-3023000301312222"></a>

## origin_servers.consul_service.outside_network — outside_network / 213321123311 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101)
- origin_servers.consul_service.outside_network

<a id="canonical-2001322200002122-1231301123332322-0300121011121113-2111202032020310-2100131213210122-0221000112331123-3302031003120303-1122102210222113"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside network.

Upstream description:

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

<a id="canonical-0331130202202332-1011001320332020-3133002011012200-2231130113301002-0122311110312101-2320221213000202-0012000222312321-2213003313312201"></a>

## Direct properties — outside_network / 213321123311 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2130210121012211-3331222001001301-2111110221312011-3231122111211010-3102121233312323-3203210203110323-3133023021213313-1021320013301122"></a>

## Next pages — outside_network / 213321123311 / 4

- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-0220232232031101-0021332122202322-0000010021303022-3120030132133221-3332012010321202-3213003213331110-3022311110231022-0030230300222112"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0121201133331112-3010203330030000-0111320022330120-0122233020200223-0211310210123031-1010303112201123-2333223130302201-2230332022300232"></a>

## origin_servers.consul_service.site_locator — site_locator / 120320213121 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101)
- origin_servers.consul_service.site_locator

<a id="canonical-2201002203020211-2222200021030213-2303312333023310-1303230031112302-2123323031022112-3020032000231233-0020323030031323-2211133311111301"></a>

Type: `"single"`. Computed.

Message defines a reference to a site or virtual site object.

Upstream description:

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

<a id="canonical-2300332210022022-2123103223202020-1233323001210210-1021220203132122-2223223233330212-3000203012332020-2323010130233102-2301202301230300"></a>

## Direct properties — site_locator / 120320213121 / 3

- [site](data-sources--origin_pool--reference--group-002.md#canonical-3201110303231310-3011013210011132-1203330220301213-3020102100210210-1100000013303333-0323223030101023-3101331232320122-2200012221211323): complete subsection reference.

- [virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-0011311000232122-2211001023101233-2330112333230321-1011003130303122-0023333100100231-1121320123123211-2331100112332302-0232012030310231): complete subsection reference.

<a id="canonical-2011000101202133-0001013030302000-0321322310121333-2323123123323022-3131030131112032-1111313332120212-3001030213103011-3132332121313013"></a>

## Next pages — site_locator / 120320213121 / 4

- [origin_servers.consul_service.site_locator.site](data-sources--origin_pool--reference--group-002.md#canonical-3201110303231310-3011013210011132-1203330220301213-3020102100210210-1100000013303333-0323223030101023-3101331232320122-2200012221211323)
- [origin_servers.consul_service.site_locator.virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-0011311000232122-2211001023101233-2330112333230321-1011003130303122-0023333100100231-1121320123123211-2331100112332302-0232012030310231)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-3201110303231310-3011013210011132-1203330220301213-3020102100210210-1100000013303333-0323223030101023-3101331232320122-2200012221211323"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3320201311133231-1110101302201221-1222302013300123-0301031033230232-1312103102203302-0001003312013212-0223013312122031-2311220130011213"></a>

## origin_servers.consul_service.site_locator.site — site / 000333330103 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101)
- [origin_servers.consul_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-0220232232031101-0021332122202322-0000010021303022-3120030132133221-3332012010321202-3213003213331110-3022311110231022-0030230300222112)
- origin_servers.consul_service.site_locator.site

<a id="canonical-1313003323320102-2303213312302230-3302110132132131-1110022102103022-1203310010303222-3230231330220332-3330203002223202-0232011311001021"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-1121303222103212-0000133001110231-3221130133212111-2213230003300311-0301303303223001-0103302001202010-3323101320331021-2011211121221020"></a>

## Direct properties — site / 000333330103 / 3

<a id="canonical-3000122233210221-0013332312311303-2001121210122023-1230331330203320-0111323012222032-3110013000332120-0103221311011303-2032122120322030"></a>

<a id="canonical-2310010130312321-1113210221113100-0333012230221103-2111120133310002-0211012233223030-1032231022030032-3312011210221320-3211000022011303"></a>

## name property — site / 000333330103 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0003010221112231-3301122210021031-2311333302322103-0223033330010021-3021221110202200-1101302222200231-0033200320333131-2011223230231030"></a>

<a id="canonical-1212313003133311-1113000313202120-0330123213322101-1233113013023221-2120132120320003-1111301310013230-0002323032003300-0020221101232300"></a>

## namespace property — site / 000333330103 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3211212202030121-3200311312212202-3232102102133030-3003213201311312-2301321112113100-2110002312012113-2003111200233312-0013002212013212"></a>

<a id="canonical-0120320222102033-2313211232032123-2313113333322302-3000033321133102-0022333131311311-3332321213013011-3212313331131231-1201112131303232"></a>

## tenant property — site / 000333330103 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3103231211012321-3102021123020331-3033022311112331-0201210032030112-3110333201102212-1331222221212131-3000011320023210-1012222003112121"></a>

## Next pages — site / 000333330103 / 7

- [origin_servers.consul_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-0220232232031101-0021332122202322-0000010021303022-3120030132133221-3332012010321202-3213003213331110-3022311110231022-0030230300222112)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-0011311000232122-2211001023101233-2330112333230321-1011003130303122-0023333100100231-1121320123123211-2331100112332302-0232012030310231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1013322200102312-2210300131110031-2001113022223223-0002011332022031-0311200311231000-3200333130110103-3310102021032112-1001011111101330"></a>

## origin_servers.consul_service.site_locator.virtual_site — virtual_site / 203112112230 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101)
- [origin_servers.consul_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-0220232232031101-0021332122202322-0000010021303022-3120030132133221-3332012010321202-3213003213331110-3022311110231022-0030230300222112)
- origin_servers.consul_service.site_locator.virtual_site

<a id="canonical-0020302332213002-0211232210010332-0210230023321103-3012221020333233-3021321021313323-2331131012123100-0111302121230001-2322121202303023"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-2231120302302331-2232311200213102-1000202212302232-1013133020303201-3223031203133303-2003230132012031-2311113301110323-3310302000013220"></a>

## Direct properties — virtual_site / 203112112230 / 3

<a id="canonical-3311231130121310-1113010133301031-3110023330122112-0011201101203001-2212121300210130-1022032332122003-1223013001021303-0331233101011330"></a>

<a id="canonical-1133331023010111-2130002101010231-1012223312102212-0220122320010332-2131310003133213-2331033002023023-2120020310031211-2333120021203000"></a>

## name property — virtual_site / 203112112230 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0323311030133302-1020031231202313-2300133333321203-3012101332012111-0101001103021222-3302310211232223-0302200110133302-3012021120303010"></a>

<a id="canonical-3303111010130100-3330320331021311-3133132230012311-1121321013220113-0311030022232121-2010331222132213-0120133112200221-3230301212311022"></a>

## namespace property — virtual_site / 203112112230 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2031011002312332-0112122212112020-2021313002121220-1231113333220302-3002100213102213-2023133130021302-1113131023321112-2031000313113103"></a>

<a id="canonical-1232113010221300-3221113222201311-2123331103021100-3031200010003133-1300221231113111-1222221222033100-1312011211321311-1211322220213023"></a>

## tenant property — virtual_site / 203112112230 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1222103130022132-1021321112231012-1200011110010111-1201223031121221-0303232303012111-2123201322102332-2220111313321103-0212133221222323"></a>

## Next pages — virtual_site / 203112112230 / 7

- [origin_servers.consul_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-0220232232031101-0021332122202322-0000010021303022-3120030132133221-3332012010321202-3213003213331110-3022311110231022-0030230300222112)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2320132102213312-1002323332032110-2303031210111010-2303200111311303-3103313332001302-2113122313133300-0233103020312222-0232102301003021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201010013122030-3111333111321221-2132123221231311-1032303203023321-2002231002230112-0330320310111111-3033033023102110-1013010101323201"></a>

## origin_servers.consul_service.snat_pool — snat_pool / 021101333020 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101)
- origin_servers.consul_service.snat_pool

<a id="canonical-2002312100030302-3002022232311220-3301132331131212-0010310320111231-1113213010131230-3023023210130222-3222233301023132-3203022211022202"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

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

<a id="canonical-3203333000022001-3111021121331330-3211130230232303-1231023001030323-2320111332313313-0201030223230311-3012220333233133-2103103013133100"></a>

## Direct properties — snat_pool / 021101333020 / 3

- [no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-2212302023010121-3012311222032022-2331132001202020-1012020032233011-2122122310311010-3312003010312122-1030011201211321-1012203301322300): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-1332111321010200-3023123230030001-0220322303300123-2032020220130332-2332222111130213-3203221213003320-0332211023323203-1220030013121033): complete subsection reference.

<a id="canonical-1110222231331300-3232113211012221-1011221021322301-2030222120333001-1220101310030103-0020233212221333-1121221002313303-3121033130322230"></a>

## Next pages — snat_pool / 021101333020 / 4

- [origin_servers.consul_service.snat_pool.no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-2212302023010121-3012311222032022-2331132001202020-1012020032233011-2122122310311010-3312003010312122-1030011201211321-1012203301322300)
- [origin_servers.consul_service.snat_pool.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-1332111321010200-3023123230030001-0220322303300123-2032020220130332-2332222111130213-3203221213003320-0332211023323203-1220030013121033)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2212302023010121-3012311222032022-2331132001202020-1012020032233011-2122122310311010-3312003010312122-1030011201211321-1012203301322300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0201123233311030-2313013200010223-2020212330300303-3332322212121333-0212103130331333-3303032110132331-1331202333203220-1033103331000221"></a>

## origin_servers.consul_service.snat_pool.no_snat_pool — no_snat_pool / 111233201202 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101)
- [origin_servers.consul_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-2320132102213312-1002323332032110-2303031210111010-2303200111311303-3103313332001302-2113122313133300-0233103020312222-0232102301003021)
- origin_servers.consul_service.snat_pool.no_snat_pool

<a id="canonical-3202201023102303-2012010023330120-3312200223102300-3321033012123312-1000320132320311-2203022332211232-3230313130120111-3031211232132110"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no snat pool.

Upstream description:

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

<a id="canonical-0021231321020303-1313230030002020-1033013302132232-0113033232113023-1001011012001321-0021032210200010-3121103300103221-3300331033310010"></a>

## Direct properties — no_snat_pool / 111233201202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2300110231102030-2323321220313321-1333002200011232-1332323011021333-1113101210330133-0011003133201220-2232302121220031-2320002323320212"></a>

## Next pages — no_snat_pool / 111233201202 / 4

- [origin_servers.consul_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-2320132102213312-1002323332032110-2303031210111010-2303200111311303-3103313332001302-2113122313133300-0233103020312222-0232102301003021)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-1332111321010200-3023123230030001-0220322303300123-2032020220130332-2332222111130213-3203221213003320-0332211023323203-1220030013121033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1310312232011011-1113113233120002-0213222132233300-0330333111331322-3122110213210300-0102113002313122-0222311200211320-1330031223100032"></a>

## origin_servers.consul_service.snat_pool.snat_pool — snat_pool / 301023231211 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.consul_service](data-sources--origin_pool--reference--group-001.md#canonical-2232200030021030-0131221020303000-0013003112130202-2121201300232031-0113011133100221-1030000313002112-3102203202200222-0112021202200101)
- [origin_servers.consul_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-2320132102213312-1002323332032110-2303031210111010-2303200111311303-3103313332001302-2113122313133300-0233103020312222-0232102301003021)
- origin_servers.consul_service.snat_pool.snat_pool

<a id="canonical-3022220313021103-1130003021313320-1220333003113330-0332112323300013-1321022310011020-0200012202331312-2003202332120302-3203310202303332"></a>

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

<a id="canonical-2022100221022222-3123310312020130-0223200122212002-0103230021130322-2010201121002033-1101313011231000-2301122323201103-1212200203102310"></a>

## Direct properties — snat_pool / 301023231211 / 3

<a id="canonical-0210132323210231-2130202320012330-1103212300013211-1303200213300003-2011130211020101-2113031321300113-2113133111232103-3113201213020222"></a>

<a id="canonical-0110222003122200-2230132302112331-2331033331001230-0230032123313330-1112330020001103-3300110100000233-2212130113232312-3020323012122032"></a>

## prefixes property — snat_pool / 301023231211 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0311200221300203-1210232023112112-2133132121013311-1131222122321332-0020003133201233-1002210031110230-3212131233312001-3300303001132103"></a>

## Next pages — snat_pool / 301023231211 / 5

- [origin_servers.consul_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-2320132102213312-1002323332032110-2303031210111010-2303200111311303-3103313332001302-2113122313133300-0233103020312222-0232102301003021)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-3123002010230120-1211300002023332-3220012210300131-1320202221223001-3220302001323203-2022021333133032-1231223202322213-0303321213212231"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213212110323120-0303110023333310-0133030300113132-0303321031011330-3131130300031213-1312212202232120-2220300221031022-3113121101122202"></a>

## origin_servers.custom_endpoint_object — custom_endpoint_object / 032231000020 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- origin_servers.custom_endpoint_object

<a id="canonical-3333021302313233-2132310320013112-0113132311331331-2123210023300020-2030301130113303-3310200033000000-3230303211313120-3220310031310203"></a>

Type: `"single"`. Computed.

Specify origin server with a reference to endpoint object.

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

<a id="canonical-2023102131333123-2020312330132231-0021030103303001-3123002112013321-0120022331033012-2331021323223221-0122033221001213-2121333210033232"></a>

## Direct properties — custom_endpoint_object / 032231000020 / 3

- [endpoint](data-sources--origin_pool--reference--group-002.md#canonical-2200222130022302-2303131210112110-0123312133022300-3033321312313320-1231231221303333-1111222311213303-1222113011223022-2131203212012131): complete subsection reference.

<a id="canonical-0001332301300122-2100231330332011-1220011101202023-2111303032310000-3122233212302222-3130033130101113-0332132102322110-0130020101032330"></a>

## Next pages — custom_endpoint_object / 032231000020 / 4

- [origin_servers.custom_endpoint_object.endpoint](data-sources--origin_pool--reference--group-002.md#canonical-2200222130022302-2303131210112110-0123312133022300-3033321312313320-1231231221303333-1111222311213303-1222113011223022-2131203212012131)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2200222130022302-2303131210112110-0123312133022300-3033321312313320-1231231221303333-1111222311213303-1222113011223022-2131203212012131"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2322111001211120-0011323303313313-0222031221011132-1130302233301131-1102113322102103-0332031110122221-2102101121221011-2123300332032032"></a>

## origin_servers.custom_endpoint_object.endpoint — endpoint / 010132110233 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.custom_endpoint_object](data-sources--origin_pool--reference--group-002.md#canonical-3123002010230120-1211300002023332-3220012210300131-1320202221223001-3220302001323203-2022021333133032-1231223202322213-0303321213212231)
- origin_servers.custom_endpoint_object.endpoint

<a id="canonical-1213110323211231-0333323132100100-0111033023230312-2120123301202333-1032110200313211-0223020233100211-2322033020203321-3333100211133201"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-1320233120322331-3002120100322131-1233011222220203-1332012223312202-0301032121310013-0122021312310101-0202022313002130-1212101130112321"></a>

## Direct properties — endpoint / 010132110233 / 3

<a id="canonical-3131322332321103-3233133102210301-3233120011202133-0002123111003220-2322230232312331-1201201301020111-0130232021332013-3232031000112323"></a>

<a id="canonical-2000122223033220-1110012203022121-1103322303302223-2103113222200112-1212100323211302-3123230321013202-3120120002000033-0032330132331120"></a>

## name property — endpoint / 010132110233 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0320122000303100-2211122300113100-1231000112232003-1100111233232010-0300033333232333-3120330230313123-2223100212332311-3320221232202122"></a>

<a id="canonical-0001232111322032-0020020200321010-2103201201001033-2102123021233031-2131313321300220-2300033220123021-1002202212323132-2223000232230001"></a>

## namespace property — endpoint / 010132110233 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3133030202000321-3301010003320331-3110322022313022-2001231320013112-2132002020322202-2301020022310311-3003201201021101-0311020122320122"></a>

<a id="canonical-0321012021332300-0030000003132020-0022122200103323-2231111010131330-1233132233233002-2101101233310023-2102322012320301-0102132331232033"></a>

## tenant property — endpoint / 010132110233 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0311230022120102-1232220011223321-2330321303121230-3201320122122032-0231031210313012-0232112212032021-2213311223303223-3011010310220121"></a>

## Next pages — endpoint / 010132110233 / 7

- [origin_servers.custom_endpoint_object](data-sources--origin_pool--reference--group-002.md#canonical-3123002010230120-1211300002023332-3220012210300131-1320202221223001-3220302001323203-2022021333133032-1231223202322213-0303321213212231)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202032000113203-3311021032200010-0020022020201310-2011320211120033-2131122131103320-0021113100001111-2220103333122112-2031032110000133"></a>

## origin_servers.k8s_service — k8s_service / 033201013032 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- origin_servers.k8s_service

<a id="canonical-1223233020022101-3220211210102003-1322331302233310-3210312321210332-2022313212330223-3032211030303222-1023120103333200-1113332132013312"></a>

Type: `"single"`. Computed.

Specify origin server with K8s service name and site information.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"inside_network\",\"outside_network\",\"vk8s_networks\"]",
  "x-ves-oneof-field-service_info": "[\"service_name\"]"
}
```

<a id="canonical-1302323302021001-2133133023303320-2210100303100202-3221203233122010-0330301210213133-2222112320121300-2330333103203122-2101221123330230"></a>

## Direct properties — k8s_service / 033201013032 / 3

- [inside_network](data-sources--origin_pool--reference--group-002.md#canonical-2121000330003222-2030211333231213-1131222130230302-3113030331232101-3122113231312313-2202202111213213-1222131313231130-3111112122202202): complete subsection reference.

- [outside_network](data-sources--origin_pool--reference--group-002.md#canonical-2010233310313303-2232131212223021-1021133333220302-0031003331120131-1103201003023210-3231233320021013-3312300231122222-3032320221011203): complete subsection reference.

<a id="canonical-2033110102321231-2011330302322202-1233110020223330-2221122203002103-0300002033001031-3320102302330233-0332002101233021-3003323032022011"></a>

<a id="canonical-2123002203100223-1130133222120033-0102120331033003-0022103133311110-1123223030112113-1220331031013102-1010200213201202-2000133013201021"></a>

## protocol property — k8s_service / 033201013032 / 4

Type: `"string"`. Computed.

\[Enum: PROTOCOL\_TCP|PROTOCOL\_UDP\] Type of protocol - PROTOCOL\_TCP: TCP - PROTOCOL\_UDP: UDP.
Possible values are \`PROTOCOL\_TCP\`, \`PROTOCOL\_UDP\`. Defaults to \`PROTOCOL\_TCP\`.

Upstream description:

Type of protocol

&#8203;- PROTOCOL\_TCP: TCP

&#8203;- PROTOCOL\_UDP: UDP.

Receipt-pinned upstream constraints:

```json
{
  "default": "PROTOCOL_TCP",
  "enum": [
    "PROTOCOL_TCP",
    "PROTOCOL_UDP"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="canonical-2203033320222001-2103023111123222-2123233202333022-1302301030303113-0032310111031133-3231203233112102-2313112330103011-0010231113212013"></a>

<a id="canonical-0300300113300230-3313221330111123-1200030202200303-0133003113230121-1123320321022101-2223331011021321-1003013122133312-1000021230222023"></a>

## service_name property — k8s_service / 033201013032 / 5

Type: `"string"`. Computed.

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace'frontend', namespace is 'speedtest' and cluster-ID is
'prod', then you will enter..

Upstream description:

Exclusive with \[\] K8s service name of the origin server will be listed, including the namespace
and cluster-ID. For vK8s services, you need to enter a string with the format
servicename.namespace:example-namespace"frontend", namespace is "speedtest" and cluster-ID is
"prod", then you will enter "frontend.speedtest:prod". Both namespace and cluster-ID are optional.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ves_service_namespace_name": "true"
  }
}
```

- [site_locator](data-sources--origin_pool--reference--group-002.md#canonical-3213302323123220-2200323333301123-0321023122332230-0200100310110120-2302312313312031-3230112131212203-3301021011100222-3000003312200123): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-1003201010222201-1101110320120012-1210031020121301-2311123310120002-3313110220322012-2200010133303101-3120221211111312-0310000310102010): complete subsection reference.

- [vk8s_networks](data-sources--origin_pool--reference--group-002.md#canonical-3003320220101232-2222122002000011-0122233212223100-0113213320301211-3321322300223222-1233131112321100-3101110202223112-2221113000223003): complete subsection reference.

<a id="canonical-0310300132022132-2023001031030202-0123302301132233-2223311113322213-1300031130320102-0022112023012132-3022110111103222-0211210032300301"></a>

## Next pages — k8s_service / 033201013032 / 6

- [origin_servers.k8s_service.inside_network](data-sources--origin_pool--reference--group-002.md#canonical-2121000330003222-2030211333231213-1131222130230302-3113030331232101-3122113231312313-2202202111213213-1222131313231130-3111112122202202)
- [origin_servers.k8s_service.outside_network](data-sources--origin_pool--reference--group-002.md#canonical-2010233310313303-2232131212223021-1021133333220302-0031003331120131-1103201003023210-3231233320021013-3312300231122222-3032320221011203)
- [origin_servers.k8s_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-3213302323123220-2200323333301123-0321023122332230-0200100310110120-2302312313312031-3230112131212203-3301021011100222-3000003312200123)
- [origin_servers.k8s_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-1003201010222201-1101110320120012-1210031020121301-2311123310120002-3313110220322012-2200010133303101-3120221211111312-0310000310102010)
- [origin_servers.k8s_service.vk8s_networks](data-sources--origin_pool--reference--group-002.md#canonical-3003320220101232-2222122002000011-0122233212223100-0113213320301211-3321322300223222-1233131112321100-3101110202223112-2221113000223003)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2121000330003222-2030211333231213-1131222130230302-3113030331232101-3122113231312313-2202202111213213-1222131313231130-3111112122202202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1011023331133231-0201000102020312-1020333113331331-3000123130131103-1333200202300213-2110332122010220-1030023303023000-2130013102223220"></a>

## origin_servers.k8s_service.inside_network — inside_network / 203010213301 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- origin_servers.k8s_service.inside_network

<a id="canonical-2203311331100022-3222230321303230-2120012031212131-1200031230202303-0301032213011323-0233003220100333-0111101122312331-1300222230302132"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside network.

Upstream description:

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

<a id="canonical-3110320200222120-2033021000222310-1112221301111313-1003303020030330-1131302212231332-0030000200300332-0102032100001230-3011023212332002"></a>

## Direct properties — inside_network / 203010213301 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2031131312131221-2311220103023300-3030121330110022-1211331322130322-1231031011000302-2032202223310203-1222022333032130-0021301122202231"></a>

## Next pages — inside_network / 203010213301 / 4

- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2010233310313303-2232131212223021-1021133333220302-0031003331120131-1103201003023210-3231233320021013-3312300231122222-3032320221011203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2103212101213300-2000310301301310-1013231023220023-2213113000222311-1213113030120310-3130123333011302-0020300101131320-2310213201110230"></a>

## origin_servers.k8s_service.outside_network — outside_network / 232223320032 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- origin_servers.k8s_service.outside_network

<a id="canonical-3012021223332031-2331031010321322-1321023001130222-1110130231311222-1030302202321020-0010220322233210-0200002000311313-0030121002132231"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside network.

Upstream description:

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

<a id="canonical-0122111022022200-3222101033331322-3001223232211322-3222130213311311-1100132212232022-2003312203333112-2010122300211122-1122220222331213"></a>

## Direct properties — outside_network / 232223320032 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3111222211223310-3231112322120132-3010020132003122-0303022213033130-1222021030122021-1310321013230023-1113312133003211-1300201101233021"></a>

## Next pages — outside_network / 232223320032 / 4

- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-3213302323123220-2200323333301123-0321023122332230-0200100310110120-2302312313312031-3230112131212203-3301021011100222-3000003312200123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0232212300132300-1000231001231303-1030131000332223-0303220212302030-3013212120102020-1131020323333130-0322033233331321-3233301332122010"></a>

## origin_servers.k8s_service.site_locator — site_locator / 220223301320 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- origin_servers.k8s_service.site_locator

<a id="canonical-0022112310330022-0110101102111300-2022230203223313-1021300231102100-1320312212011221-0220311212012120-0021321021212233-1122023130231201"></a>

Type: `"single"`. Computed.

Message defines a reference to a site or virtual site object.

Upstream description:

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

<a id="canonical-0232320200123313-0022232331313103-0302203001231303-0212233022100310-0001100132311321-3000030132132100-2000232220103322-2300121202221022"></a>

## Direct properties — site_locator / 220223301320 / 3

- [site](data-sources--origin_pool--reference--group-002.md#canonical-0023122133213320-0032120210322223-1233002230031131-1320321120020013-2032223202303301-1220121113011213-2120311220321203-3212310020202210): complete subsection reference.

- [virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-3032123210002103-2121113021330010-0033023103002023-0230300130130333-2031031021203312-0133011211103322-3002120330033330-3132122121131022): complete subsection reference.

<a id="canonical-2011202230022230-1003122222130010-0210300230101001-1203011120131020-2321130232112132-2113032313330101-0010221210001022-3012011332301012"></a>

## Next pages — site_locator / 220223301320 / 4

- [origin_servers.k8s_service.site_locator.site](data-sources--origin_pool--reference--group-002.md#canonical-0023122133213320-0032120210322223-1233002230031131-1320321120020013-2032223202303301-1220121113011213-2120311220321203-3212310020202210)
- [origin_servers.k8s_service.site_locator.virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-3032123210002103-2121113021330010-0033023103002023-0230300130130333-2031031021203312-0133011211103322-3002120330033330-3132122121131022)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-0023122133213320-0032120210322223-1233002230031131-1320321120020013-2032223202303301-1220121113011213-2120311220321203-3212310020202210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1300101013110220-0322232031130322-1002321021220202-1320033000102301-0201323302103002-1231113300310102-0003030331230023-2300123321000130"></a>

## origin_servers.k8s_service.site_locator.site — site / 331303201213 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- [origin_servers.k8s_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-3213302323123220-2200323333301123-0321023122332230-0200100310110120-2302312313312031-3230112131212203-3301021011100222-3000003312200123)
- origin_servers.k8s_service.site_locator.site

<a id="canonical-1132230003030301-1031133202102201-3111000112233130-2111001033320033-0323131211130233-2231310022112003-0220312223332322-0203131100022311"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-2231022101211132-1131020311231020-1312130302132301-3032201202123212-0230323001303032-0013013011003111-2232132231210102-0110121130031313"></a>

## Direct properties — site / 331303201213 / 3

<a id="canonical-2020201212221111-3011330313010203-2021202000321221-1103133213210120-1123021302330132-0030122133300010-1112113130033231-0132212231001120"></a>

<a id="canonical-1321211131313112-0323120320231000-2133302231131300-2131333133311021-1102300310131220-1331232322102002-3103313333210022-3023030010303203"></a>

## name property — site / 331303201213 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0202233210233213-0013230131322323-0121011223032222-3012120221000302-0120002202022021-0130001311111213-0123123020123322-1200003023220321"></a>

<a id="canonical-2211302303302312-2223032233123301-3213110330212103-3213112130003110-3010313223010230-0200200013011031-0213220222122121-2222212002312011"></a>

## namespace property — site / 331303201213 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2222012313310210-0022032013100010-1211230330120201-1111302301322301-3113122201331221-3001023023003210-1031121200020003-1131223333333211"></a>

<a id="canonical-1330300001031023-3022001233120113-1332110131032332-1122222111332022-0111130021100220-2230023230202121-3213110302012131-3030101230011100"></a>

## tenant property — site / 331303201213 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2212222113230011-2010011332203103-3120133230033311-1012001001120213-0001223113303023-2220200330111101-0110010101103201-0133222221332033"></a>

## Next pages — site / 331303201213 / 7

- [origin_servers.k8s_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-3213302323123220-2200323333301123-0321023122332230-0200100310110120-2302312313312031-3230112131212203-3301021011100222-3000003312200123)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-3032123210002103-2121113021330010-0033023103002023-0230300130130333-2031031021203312-0133011211103322-3002120330033330-3132122121131022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232100330223100-0321100301330200-1030211312313011-3123203032003100-2212211230302103-1202320003131301-2012132213332133-2100031011333130"></a>

## origin_servers.k8s_service.site_locator.virtual_site — virtual_site / 313231212223 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- [origin_servers.k8s_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-3213302323123220-2200323333301123-0321023122332230-0200100310110120-2302312313312031-3230112131212203-3301021011100222-3000003312200123)
- origin_servers.k8s_service.site_locator.virtual_site

<a id="canonical-0310032322220213-0023320311333122-3300011313031132-3332003333030021-3202203330303110-1320232211123233-2113232133332122-2302112132103200"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-3312011321001100-3320123303202211-1201223331112313-2021131122222212-2020103133123202-0211110031310122-0312222113202112-1113223312110123"></a>

## Direct properties — virtual_site / 313231212223 / 3

<a id="canonical-3121121032012310-3323221103112303-3033102120210112-1312011301331103-1011030030011113-3313210023300001-3311121031021013-3333031222000131"></a>

<a id="canonical-0002333032300210-2321320223002001-2001223110232333-1023213033132220-1111003112312230-1030131002222211-2130213000103331-3212333012213113"></a>

## name property — virtual_site / 313231212223 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3100000122100212-0311032300110231-2102310101333300-3101201323130231-2330121223111103-2020301113203031-1122031231022020-1213022011303232"></a>

<a id="canonical-0133300020032303-3003021333001100-2001233311102203-1330221312100221-0003121312102200-3031133330310122-3202102320033233-3121333121111000"></a>

## namespace property — virtual_site / 313231212223 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3332320101300202-3201001213312032-3303023212202202-1132003110303221-0212331033232323-3001120123302013-1131200003231103-2320233302233320"></a>

## tenant property — virtual_site / 313231212223 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3032203201211020-1323322311132023-2332213330212133-3203033002003230-3013010023231001-2311312001303321-1200213012032101-2331212201232020"></a>

## Next pages — virtual_site / 313231212223 / 7

- [origin_servers.k8s_service.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-3213302323123220-2200323333301123-0321023122332230-0200100310110120-2302312313312031-3230112131212203-3301021011100222-3000003312200123)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-1003201010222201-1101110320120012-1210031020121301-2311123310120002-3313110220322012-2200010133303101-3120221211111312-0310000310102010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302000213203122-2322302201010312-3131020110113132-0112030121330102-3210120200001000-0310023231022100-1021223112320332-2233002030300011"></a>

## origin_servers.k8s_service.snat_pool — snat_pool / 331211011312 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- origin_servers.k8s_service.snat_pool

<a id="canonical-2300132112011133-1001222102231111-1021033011211330-1013213202333130-1031330203222122-1031221000101110-0232123121022112-2111300211032002"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

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

<a id="canonical-1222032100101230-3301130323012103-3203223201331232-3202003031302123-2120223133222313-2110302311020303-0110133022023333-0320221321203112"></a>

## Direct properties — snat_pool / 331211011312 / 3

- [no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-2021331300301000-0201112021010013-3030120000311113-3213033221203332-1322203033021022-0000202221101311-2200013212221201-1311010002330321): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-1331300313220011-3023330120132021-3001001101013210-3130003312112001-1122121110300022-3203003032232113-2301002111100331-2233032032111230): complete subsection reference.

<a id="canonical-1200122333310323-2302113203313212-3203220312000021-3020332222111210-1210113002333302-3131033221220200-1313103330233320-1321033123220030"></a>

## Next pages — snat_pool / 331211011312 / 4

- [origin_servers.k8s_service.snat_pool.no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-2021331300301000-0201112021010013-3030120000311113-3213033221203332-1322203033021022-0000202221101311-2200013212221201-1311010002330321)
- [origin_servers.k8s_service.snat_pool.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-1331300313220011-3023330120132021-3001001101013210-3130003312112001-1122121110300022-3203003032232113-2301002111100331-2233032032111230)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2021331300301000-0201112021010013-3030120000311113-3213033221203332-1322203033021022-0000202221101311-2200013212221201-1311010002330321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3232032332113003-3011222203003221-3310120113300121-0030200222001203-3120022323033001-1232130310202032-0322001303101320-1013121211012223"></a>

## origin_servers.k8s_service.snat_pool.no_snat_pool — no_snat_pool / 011300211231 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- [origin_servers.k8s_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-1003201010222201-1101110320120012-1210031020121301-2311123310120002-3313110220322012-2200010133303101-3120221211111312-0310000310102010)
- origin_servers.k8s_service.snat_pool.no_snat_pool

<a id="canonical-2112133302313122-0101020213222322-2101220231122310-1330323112021031-2030200331033020-1031013313100001-2013232120021223-0320212020313322"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no snat pool.

Upstream description:

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

<a id="canonical-1212332100131313-1120333131131000-1203123232020203-2332123120120012-3001101231312302-1031310002122211-2301231223212030-3030133210302001"></a>

## Direct properties — no_snat_pool / 011300211231 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-3002023103220323-2213110132103310-0012301130000020-1231310333012333-3232100333221320-3012200211030212-3322033030200132-3312130321211223"></a>

## Next pages — no_snat_pool / 011300211231 / 4

- [origin_servers.k8s_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-1003201010222201-1101110320120012-1210031020121301-2311123310120002-3313110220322012-2200010133303101-3120221211111312-0310000310102010)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-1331300313220011-3023330120132021-3001001101013210-3130003312112001-1122121110300022-3203003032232113-2301002111100331-2233032032111230"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2102020330201322-0102303331123001-3331100200123323-3020311003210332-0122212213132133-2121123231330003-2233302021000231-2113231030122221"></a>

## origin_servers.k8s_service.snat_pool.snat_pool — snat_pool / 203031020221 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- [origin_servers.k8s_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-1003201010222201-1101110320120012-1210031020121301-2311123310120002-3313110220322012-2200010133303101-3120221211111312-0310000310102010)
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

<a id="canonical-2323103122233310-0101302331103221-3013033311002310-2021320122100101-1101310221302002-0022100332013002-1120021320301133-1333330321220130"></a>

## Direct properties — snat_pool / 203031020221 / 3

<a id="canonical-1200301233210100-3313112203023312-2100103121232311-2202010012221101-2330320031030010-3031220212201302-3132221301132002-1030022001210021"></a>

<a id="canonical-0201030302100133-0311300122120032-3201330212221330-3211113221032211-0103031121003031-0321330311012021-3312230012330100-0222133130312013"></a>

## prefixes property — snat_pool / 203031020221 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0101220121213233-1112113020312200-2021020112311230-1303132311200203-0033113321303200-0210332210302300-0232310001211131-1122113232000111"></a>

## Next pages — snat_pool / 203031020221 / 5

- [origin_servers.k8s_service.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-1003201010222201-1101110320120012-1210031020121301-2311123310120002-3313110220322012-2200010133303101-3120221211111312-0310000310102010)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-3003320220101232-2222122002000011-0122233212223100-0113213320301211-3321322300223222-1233131112321100-3101110202223112-2221113000223003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0302111033010013-2111122332211312-2213223201323101-2322210311312220-2102233120311132-1101220023132010-3210300312323300-0021111003000002"></a>

## origin_servers.k8s_service.vk8s_networks — vk8s_networks / 112021322202 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- origin_servers.k8s_service.vk8s_networks

<a id="canonical-1121320221330002-3130133132310313-3123111112030121-2200020223033312-3302202200233301-1203100203000022-2113333233111312-0022330210032111"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for vk8s networks.

Upstream description:

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

<a id="canonical-2031023012000000-2221020110202022-2231032210030231-0030210032331133-2133131322301222-3201330332301311-0202301311102302-1222133312113313"></a>

## Direct properties — vk8s_networks / 112021322202 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0322103313132333-1010312110030031-0003111230231330-1201031330021030-3311011311332312-1131100203022130-2221023000320331-3010132130010303"></a>

## Next pages — vk8s_networks / 112021322202 / 4

- [origin_servers.k8s_service](data-sources--origin_pool--reference--group-002.md#canonical-2300210011321131-0121020312033232-0121113020120121-3022120202323312-0223013310333023-2001013231310202-0002322000012212-2013232333323201)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0131000333221033-2223311223311001-0302133323312210-1333030210203222-2231323330022221-1030022220213331-1200112100100032-3333203022330310"></a>

## origin_servers.private_ip — private_ip / 323200220213 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
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

<a id="canonical-3201122320201100-2232003003132312-1302121322103210-0213323033333222-3121122113032011-1020130230302311-3223031113311000-0333322103003223"></a>

## Direct properties — private_ip / 323200220213 / 3

- [inside_network](data-sources--origin_pool--reference--group-002.md#canonical-1023322133300312-3212322320031221-3000023221010020-2322001122012021-2011020320000233-1313013221012320-2120101222312123-0030101330210022): complete subsection reference.

<a id="canonical-0320121111021001-2203023121330110-3022120020120210-2221003230002312-3113001022013212-2213122121300300-1322013112110201-3332221331003322"></a>

<a id="canonical-2002210321200331-1300122230111213-2323330133322111-2221302032210100-1001113300122321-0133212200022122-1201223331113122-0223121210313133"></a>

## ip property — private_ip / 323200220213 / 4

Type: `"string"`. Computed.

IP. Exclusive with \[\] Private IPv4 address.

Upstream description:

Exclusive with \[\] Private IPv4 address.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [outside_network](data-sources--origin_pool--reference--group-002.md#canonical-0302202131230001-0233123210023202-0133201202333100-2322210001221331-3212233002323211-0121012333222123-2201123232222110-3213201123133132): complete subsection reference.

- [segment](data-sources--origin_pool--reference--group-002.md#canonical-3132303123311030-2330310310330102-0101200222301031-3110233301111120-0233230013023010-0223121320310230-2112211212011300-1103020210122222): complete subsection reference.

- [site_locator](data-sources--origin_pool--reference--group-002.md#canonical-1230220321001223-0010023132100120-0123100223321110-3121021330010013-0322133233101110-1121023020112003-2120101221002232-1022030000212030): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-0333322312323021-0312203002230202-3222301233233110-0330230100212001-1311203121121230-1212210203320003-3133230201312303-3231031023313001): complete subsection reference.

<a id="canonical-1021111013033103-2123203211200122-3013020012132132-0300011012122203-0130310323323111-1011230032230132-3001100101301302-2133031222321110"></a>

## Next pages — private_ip / 323200220213 / 5

- [origin_servers.private_ip.inside_network](data-sources--origin_pool--reference--group-002.md#canonical-1023322133300312-3212322320031221-3000023221010020-2322001122012021-2011020320000233-1313013221012320-2120101222312123-0030101330210022)
- [origin_servers.private_ip.outside_network](data-sources--origin_pool--reference--group-002.md#canonical-0302202131230001-0233123210023202-0133201202333100-2322210001221331-3212233002323211-0121012333222123-2201123232222110-3213201123133132)
- [origin_servers.private_ip.segment](data-sources--origin_pool--reference--group-002.md#canonical-3132303123311030-2330310310330102-0101200222301031-3110233301111120-0233230013023010-0223121320310230-2112211212011300-1103020210122222)
- [origin_servers.private_ip.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-1230220321001223-0010023132100120-0123100223321110-3121021330010013-0322133233101110-1121023020112003-2120101221002232-1022030000212030)
- [origin_servers.private_ip.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-0333322312323021-0312203002230202-3222301233233110-0330230100212001-1311203121121230-1212210203320003-3133230201312303-3231031023313001)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-1023322133300312-3212322320031221-3000023221010020-2322001122012021-2011020320000233-1313013221012320-2120101222312123-0030101330210022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0022322123232110-0331233202121321-2112311130212321-2220002001233203-3300031331203103-0201211302313320-0331022012111133-1031313122123023"></a>

## origin_servers.private_ip.inside_network — inside_network / 030330132200 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- origin_servers.private_ip.inside_network

<a id="canonical-1332333132310022-3003132113211323-0322120123232001-1331322020021023-3322100031003020-3013333211202122-3030113233110100-1321000310023030"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside network.

Upstream description:

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

<a id="canonical-2002101310310130-3103120231323232-2220213300213203-2323200011103020-0001010022023031-0311003320112121-2320030013213020-1133013021103021"></a>

## Direct properties — inside_network / 030330132200 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1333103122133223-2030113230311102-1111021212132312-2021231331033003-3210032013210002-3033222131333112-0030021132132111-1131032013131123"></a>

## Next pages — inside_network / 030330132200 / 4

- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-0302202131230001-0233123210023202-0133201202333100-2322210001221331-3212233002323211-0121012333222123-2201123232222110-3213201123133132"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221203131023112-2130313333123211-0321113312202123-3220031212222332-3201023110020311-2212032300003323-0200302213332230-0211103302210333"></a>

## origin_servers.private_ip.outside_network — outside_network / 222231100230 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- origin_servers.private_ip.outside_network

<a id="canonical-3031033110331132-3203320200310102-3330333312113313-0203111020203111-0102332121203313-0001011113021332-3132111312320230-3333102001033230"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside network.

Upstream description:

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

<a id="canonical-3021231121010000-0112133123021201-2201012223310131-3020131103011202-2301121302312102-0312331301101022-2301110022000030-1123122111011313"></a>

## Direct properties — outside_network / 222231100230 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1233100203332110-2321031321111112-3212001022211003-2101111102032230-2020002002230133-3110213310133210-3121101002310223-0220232210203220"></a>

## Next pages — outside_network / 222231100230 / 4

- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-3132303123311030-2330310310330102-0101200222301031-3110233301111120-0233230013023010-0223121320310230-2112211212011300-1103020210122222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130113012121210-2201002231312211-2220002010223302-1020300213230230-3200101330111112-3230021223113103-1033133033230120-0013123201302031"></a>

## origin_servers.private_ip.segment — segment / 121033210002 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- origin_servers.private_ip.segment

<a id="canonical-3201111022301122-3212230121302012-1222322302333010-2301231110023000-2310203230000220-3201031222233121-1110110313332120-2123331211011100"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-1231301232000220-3333303332321031-0002022301210111-1220221031031000-0313110233211110-3330222013200311-1001023222032333-2112102111012020"></a>

## Direct properties — segment / 121033210002 / 3

<a id="canonical-1121332010322010-0331001213013201-2131310002211300-2222320222312010-2102333233133131-1130310232202313-0022013011121321-0310220023123313"></a>

<a id="canonical-2023330212202102-3122031020130003-2211100102101231-3111223012221202-3321321130123233-2220310031322110-3000002023123112-1111311330233122"></a>

## name property — segment / 121033210002 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0213203310221310-3321113321321032-3112300331030011-3201110203011202-3201122212313012-1120002201012302-3002030110102210-1123300320320100"></a>

## namespace property — segment / 121033210002 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0131032023032301-1011022101002302-2011010300010223-3100220301031032-2313110322202001-3011021223203023-2200322131313300-0103300000333313"></a>

## tenant property — segment / 121033210002 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0232033212310301-2011101003301310-1101013231031002-3101032321132023-2332303223111011-2230131101110320-2022112132223222-3133101023022000"></a>

## Next pages — segment / 121033210002 / 7

- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-1230220321001223-0010023132100120-0123100223321110-3121021330010013-0322133233101110-1121023020112003-2120101221002232-1022030000212030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2101123010023032-1213333020121123-2032022222030201-1131113333303023-2111010111011011-3002313131223221-2313003012102120-2132120110112203"></a>

## origin_servers.private_ip.site_locator — site_locator / 222222133201 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- origin_servers.private_ip.site_locator

<a id="canonical-1122123131011031-0313203023211013-3301232131102222-0032133132333001-3213212231100211-0032200303000013-2221221320000103-0130032103113303"></a>

Type: `"single"`. Computed.

Message defines a reference to a site or virtual site object.

Upstream description:

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

<a id="canonical-0033000322010230-0310331213300003-1120103201233011-2210102301223113-1212321121323021-0013233102220131-3013210110312220-3232312302220000"></a>

## Direct properties — site_locator / 222222133201 / 3

- [site](data-sources--origin_pool--reference--group-002.md#canonical-1222313013302012-2300030200323111-3003211100300122-2203322122311323-2300302211112100-0001033202331322-2133203310023002-2101221311103322): complete subsection reference.

- [virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-3010012311121222-3323321022012103-2213010123123311-0101103202010312-1332101333222321-3332100203330302-1302312132133132-1230033300301223): complete subsection reference.

<a id="canonical-1212310120001301-1122210321302110-2123012023033101-1210033011331210-1213323001022201-1301321023032002-0301322130232031-2020303211103320"></a>

## Next pages — site_locator / 222222133201 / 4

- [origin_servers.private_ip.site_locator.site](data-sources--origin_pool--reference--group-002.md#canonical-1222313013302012-2300030200323111-3003211100300122-2203322122311323-2300302211112100-0001033202331322-2133203310023002-2101221311103322)
- [origin_servers.private_ip.site_locator.virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-3010012311121222-3323321022012103-2213010123123311-0101103202010312-1332101333222321-3332100203330302-1302312132133132-1230033300301223)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-1222313013302012-2300030200323111-3003211100300122-2203322122311323-2300302211112100-0001033202331322-2133203310023002-2101221311103322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2113032330203103-3022033332230220-3033321231332130-0203232030131330-2203231012123220-2123132131101131-0201132122113320-2122010330132100"></a>

## origin_servers.private_ip.site_locator.site — site / 010021223030 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- [origin_servers.private_ip.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-1230220321001223-0010023132100120-0123100223321110-3121021330010013-0322133233101110-1121023020112003-2120101221002232-1022030000212030)
- origin_servers.private_ip.site_locator.site

<a id="canonical-0210110200112220-0223300031210231-3200011201331330-0013221232313333-3331130210310133-3320221312023331-0301221002100220-1112123322000300"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-1033210032023010-2201313323020200-3021320220023021-2222323113212230-3212131111323110-3310133023231221-0130111203313302-2000212221222111"></a>

## Direct properties — site / 010021223030 / 3

<a id="canonical-2011333133310033-3111112013320011-3311200323000133-0010313130202330-2012112031033000-1110111312320022-1222221033022123-1112023023123102"></a>

<a id="canonical-1031201100021313-1202213331112132-2103112211211102-2321330212202320-3031231202201210-1011132120123021-2212021313231022-2130100003020233"></a>

## name property — site / 010021223030 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2231221103313123-0230101013120320-1033323031101000-3321212012321333-1233322233011200-3032210033210211-2113233322000310-2110203120300213"></a>

## namespace property — site / 010021223030 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2011202032202310-0312331031321203-2101332121012123-2120313321010100-1121232031020111-0112103022012111-0013313231101303-0202302210321220"></a>

## tenant property — site / 010021223030 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3133200001301010-1133312112100101-1022111310231311-1030211110102131-1033300231332101-0133200333322223-3113020330303013-2122301112322020"></a>

## Next pages — site / 010021223030 / 7

- [origin_servers.private_ip.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-1230220321001223-0010023132100120-0123100223321110-3121021330010013-0322133233101110-1121023020112003-2120101221002232-1022030000212030)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-3010012311121222-3323321022012103-2213010123123311-0101103202010312-1332101333222321-3332100203330302-1302312132133132-1230033300301223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2310310230102303-3101320132323313-0321002031213213-0322100203331203-2103123313000300-2230101210131221-0031212111211310-3010203103212221"></a>

## origin_servers.private_ip.site_locator.virtual_site — virtual_site / 200321030110 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- [origin_servers.private_ip.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-1230220321001223-0010023132100120-0123100223321110-3121021330010013-0322133233101110-1121023020112003-2120101221002232-1022030000212030)
- origin_servers.private_ip.site_locator.virtual_site

<a id="canonical-2322223312023130-3220100111322322-2331222033031323-2020102013312300-1021033311223033-1331102312103133-0121121111322321-0311300313201102"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-2222301011330011-2000300223302330-1133220231301120-2222111311223132-0313131312122012-1320203102321033-0232323233032322-0022211231013021"></a>

## Direct properties — virtual_site / 200321030110 / 3

<a id="canonical-2122220122323033-3011313232113110-1312231103032221-3021110310220201-0203022032302212-0313101011012210-3311303333111112-2111311031003330"></a>

<a id="canonical-2310231220122010-0003202132001330-0002203212112212-0132120021312021-2332302030030021-2011030033100232-2013211110113320-3323002022220322"></a>

## name property — virtual_site / 200321030110 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1330302210301322-0001320221010202-2030010331110210-1100210302133202-2223333022322103-3033022101130103-1010021301311321-0202101323213300"></a>

## namespace property — virtual_site / 200321030110 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1331201001332212-0133120233133112-3333310020012100-0300003233010222-0123021202013102-0302212332010300-3020112230123311-0121002322121323"></a>

## tenant property — virtual_site / 200321030110 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2332020212220102-1210311330100223-3312210313112000-0001022000133121-1212233311011312-2311200100231203-0330230201203231-0131202321201201"></a>

## Next pages — virtual_site / 200321030110 / 7

- [origin_servers.private_ip.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-1230220321001223-0010023132100120-0123100223321110-3121021330010013-0322133233101110-1121023020112003-2120101221002232-1022030000212030)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-0333322312323021-0312203002230202-3222301233233110-0330230100212001-1311203121121230-1212210203320003-3133230201312303-3231031023313001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1202000320022012-0221120302312322-3020211100110111-2212210123303210-2221123221121310-3020212112231020-0322301111302311-3023330232221113"></a>

## origin_servers.private_ip.snat_pool — snat_pool / 033003122013 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- origin_servers.private_ip.snat_pool

<a id="canonical-1230021313222323-3120023031131200-1233113301020201-2232222010011223-0032021312332211-3100132312300121-2202330322301313-1303211031130322"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

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

<a id="canonical-0020111321130012-3203302033312201-1023133323020020-2111211112002300-0133312230212123-0010131311300122-0001302213011110-0033001331011011"></a>

## Direct properties — snat_pool / 033003122013 / 3

- [no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-0232130210221013-1312312111333323-1112123110200123-0221202303332121-0121201232331122-0013121023131022-3201122223222003-1202103113231310): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-3021323101120213-2322100033001030-2033332011201330-3221222010123210-2102132032131311-0232023301002122-2021123311201013-3032110201211220): complete subsection reference.

<a id="canonical-3202020100032233-3213010133203221-2132310311220303-0232322101313323-3302231311211011-3012122103110002-0120300200031203-1320202031022331"></a>

## Next pages — snat_pool / 033003122013 / 4

- [origin_servers.private_ip.snat_pool.no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-0232130210221013-1312312111333323-1112123110200123-0221202303332121-0121201232331122-0013121023131022-3201122223222003-1202103113231310)
- [origin_servers.private_ip.snat_pool.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-3021323101120213-2322100033001030-2033332011201330-3221222010123210-2102132032131311-0232023301002122-2021123311201013-3032110201211220)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-0232130210221013-1312312111333323-1112123110200123-0221202303332121-0121201232331122-0013121023131022-3201122223222003-1202103113231310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2130021032102100-3121112021021013-3322223031330323-3113322322320223-1132031102203131-0013103301200213-0210010200212233-3302230031001122"></a>

## origin_servers.private_ip.snat_pool.no_snat_pool — no_snat_pool / 000320120322 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- [origin_servers.private_ip.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-0333322312323021-0312203002230202-3222301233233110-0330230100212001-1311203121121230-1212210203320003-3133230201312303-3231031023313001)
- origin_servers.private_ip.snat_pool.no_snat_pool

<a id="canonical-0003302321133221-3130232212332103-1122210001002331-0010132303311003-1020030123022012-0012330002212132-2232113131203303-2021031103210122"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no snat pool.

Upstream description:

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

<a id="canonical-0131233001021000-2021213221131222-0011002130001103-2333331132212302-1213312311113001-0301020322332120-1003233032322212-2320333123333031"></a>

## Direct properties — no_snat_pool / 000320120322 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1012001230033013-3003122130020313-1330003133312223-3210011130230033-0331110100330320-2331013232021102-1203131231131222-3120200023031202"></a>

## Next pages — no_snat_pool / 000320120322 / 4

- [origin_servers.private_ip.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-0333322312323021-0312203002230202-3222301233233110-0330230100212001-1311203121121230-1212210203320003-3133230201312303-3231031023313001)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-3021323101120213-2322100033001030-2033332011201330-3221222010123210-2102132032131311-0232023301002122-2021123311201013-3032110201211220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3202103031000303-3302213232000101-1211112202321223-3010210223331310-1013231330301113-1030300011201323-2021320200223220-3120222330002311"></a>

## origin_servers.private_ip.snat_pool.snat_pool — snat_pool / 301301323222 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_ip](data-sources--origin_pool--reference--group-002.md#canonical-1100302310111111-3100332123330002-1021231103102012-2031000233003021-2300001223221120-3123133220012300-2132220322302020-1322201103231102)
- [origin_servers.private_ip.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-0333322312323021-0312203002230202-3222301233233110-0330230100212001-1311203121121230-1212210203320003-3133230201312303-3231031023313001)
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

<a id="canonical-3212202020130200-2221112100301122-1210312001110010-1031122021031132-0032202023301033-1200121302010201-0120303220310111-3103302232000233"></a>

## Direct properties — snat_pool / 301301323222 / 3

<a id="canonical-2223130122312001-3221130033133232-3002313333331210-2202010202231213-0300330220003120-2310133320022001-1202230123122322-2301032302202013"></a>

<a id="canonical-0233212112211130-2032130301331220-1212322002323130-1202321030110232-1312230220302331-3201330020211132-3013132011200000-2032123302313110"></a>

## prefixes property — snat_pool / 301301323222 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2031121032321120-3113102003312131-3102311313222100-0301233030320311-2122033202110302-1102101310333220-1112332302101131-2001201323000130"></a>

## Next pages — snat_pool / 301301323222 / 5

- [origin_servers.private_ip.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-0333322312323021-0312203002230202-3222301233233110-0330230100212001-1311203121121230-1212210203320003-3133230201312303-3231031023313001)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0203000131123132-0112111212320010-1213330202100103-1113223131013332-1210123233211322-1200121001001021-1321121033223022-1230010101012010"></a>

## origin_servers.private_name — private_name / 000332232300 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
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

<a id="canonical-1131123131021030-1032110313013232-2220223221223213-0231330131110222-0312312020203132-0113202232013031-1212213323323233-1123102222223031"></a>

## Direct properties — private_name / 000332232300 / 3

<a id="canonical-0033110221311322-0232002212203030-0222030202220121-0233103311122200-2321020110302131-0120212321323311-3003313013211121-1110232232323213"></a>

<a id="canonical-2220300231003333-0220133031132210-2331130211000302-0332013302101121-2121300220301101-0112122211031302-3131030012120002-0203232200110201"></a>

## dns_name property — private_name / 000332232300 / 4

Type: `"string"`. Computed.

DNS Name. DNS Name

Upstream description:

DNS Name

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [inside_network](data-sources--origin_pool--reference--group-002.md#canonical-2201113023111301-2130033123212030-3101303112332102-0120230011113310-2333311132231210-3112223303322003-1233102300113011-1012011010000313): complete subsection reference.

- [outside_network](data-sources--origin_pool--reference--group-002.md#canonical-2333001322321030-1223130232211212-2231131122201102-2330212123111332-0233011020220031-1020133300202303-1231200222232112-1121230103123000): complete subsection reference.

<a id="canonical-0010301302220223-0122112222213313-1312231113030322-0333132012223023-3011222311030123-0201203210113311-1022202231103030-2102203101220321"></a>

<a id="canonical-2001222211312230-0011303113321303-1302123120231133-2213302212222302-2223322221021210-3000203332212230-2103330033223110-2302221230103323"></a>

## refresh_interval property — private_name / 000332232300 / 5

Type: `"number"`. Computed.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [segment](data-sources--origin_pool--reference--group-002.md#canonical-2102212203301223-1333332333313011-3203313003022320-3220031002101131-2022031112111233-0123121103301132-3023002031131212-3100331232111010): complete subsection reference.

- [site_locator](data-sources--origin_pool--reference--group-002.md#canonical-2110032230130213-1233131020100213-0103031331322233-3132300222023320-1221130110230312-3122233023110122-0213032231121113-3000120212301203): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-3312011332131233-1230133131313010-3013003121231323-0223212213120012-0301213303101202-1000312320032330-3210122031212113-2000030222032011): complete subsection reference.

<a id="canonical-3320102330313103-1333310123033033-3223111023103200-2200202230212013-1130123200021103-3102001232201130-2232312033321201-0000301100232223"></a>

## Next pages — private_name / 000332232300 / 6

- [origin_servers.private_name.inside_network](data-sources--origin_pool--reference--group-002.md#canonical-2201113023111301-2130033123212030-3101303112332102-0120230011113310-2333311132231210-3112223303322003-1233102300113011-1012011010000313)
- [origin_servers.private_name.outside_network](data-sources--origin_pool--reference--group-002.md#canonical-2333001322321030-1223130232211212-2231131122201102-2330212123111332-0233011020220031-1020133300202303-1231200222232112-1121230103123000)
- [origin_servers.private_name.segment](data-sources--origin_pool--reference--group-002.md#canonical-2102212203301223-1333332333313011-3203313003022320-3220031002101131-2022031112111233-0123121103301132-3023002031131212-3100331232111010)
- [origin_servers.private_name.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-2110032230130213-1233131020100213-0103031331322233-3132300222023320-1221130110230312-3122233023110122-0213032231121113-3000120212301203)
- [origin_servers.private_name.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-3312011332131233-1230133131313010-3013003121231323-0223212213120012-0301213303101202-1000312320032330-3210122031212113-2000030222032011)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2201113023111301-2130033123212030-3101303112332102-0120230011113310-2333311132231210-3112223303322003-1233102300113011-1012011010000313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233011100312320-3103112001012213-0033300103111311-2301311231012311-3012103113131003-2133313100011312-1213212030100121-1023302333031130"></a>

## origin_servers.private_name.inside_network — inside_network / 323000130020 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- origin_servers.private_name.inside_network

<a id="canonical-3312013330201201-3132311111230322-2232223022212333-1013203331200012-3033300123203022-0000032033020033-2030333301231201-1330312212233300"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for inside network.

Upstream description:

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

<a id="canonical-0332101110010323-1112110332133203-2020310120111223-0333231130233113-3012101222301321-3130013132310122-2302101302301203-3120033332131231"></a>

## Direct properties — inside_network / 323000130020 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1002123320110101-2112213130333110-2121121200223322-0312320011030310-1132200020303111-1102003210120230-3302331321321323-1313301031100322"></a>

## Next pages — inside_network / 323000130020 / 4

- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2333001322321030-1223130232211212-2231131122201102-2330212123111332-0233011020220031-1020133300202303-1231200222232112-1121230103123000"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2221103100121100-2222330000102002-2222212000113030-1002023120320022-0301303202023110-2203110203013133-3331203223130031-2133330220101102"></a>

## origin_servers.private_name.outside_network — outside_network / 021013221300 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- origin_servers.private_name.outside_network

<a id="canonical-0113310102130113-3100201313013101-2030200220102110-3133110031312120-0122202002021032-1033001331120103-3012332133022000-2321311332323123"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for outside network.

Upstream description:

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

<a id="canonical-0302221020012103-2021312201322223-1120033310210311-1013221003232132-2130301332123321-2002312133102103-1002133313213030-2333313100021110"></a>

## Direct properties — outside_network / 021013221300 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-0313113003312210-1332202231312323-3201200212031023-3200312021100102-2131332132310131-0332120221010301-0033330103221103-1333323311202230"></a>

## Next pages — outside_network / 021013221300 / 4

- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2102212203301223-1333332333313011-3203313003022320-3220031002101131-2022031112111233-0123121103301132-3023002031131212-3100331232111010"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2202103222301300-3312332302330022-0322223201231222-1201130123003123-0022322012003012-3013323033211213-1022331330100132-1333301302300122"></a>

## origin_servers.private_name.segment — segment / 220332312030 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- origin_servers.private_name.segment

<a id="canonical-1221001011021022-0033011011130002-2021222002011013-3120212333000223-1003102102031003-2130213321103222-2300300011213110-0222103112022322"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-2313001333100322-2012011123010210-1023221332013302-1103113311302232-1300123032021222-0013332102032033-1313211203003011-0120033331032212"></a>

## Direct properties — segment / 220332312030 / 3

<a id="canonical-1033202133100110-1322011323001330-3121210233103123-3000223222300302-0022212132321211-3301003002033303-0331130131221200-2100313113312010"></a>

<a id="canonical-0223322313332132-0023230313222001-0000223330212000-2232102203132302-3111030212201211-1103312100133021-3010203112003200-2220203330002211"></a>

## name property — segment / 220332312030 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0301123220112212-3231123031232102-2101230201020303-1032303311233003-1313303202323000-3023122200000303-2300200210100101-1312131000120000"></a>

## namespace property — segment / 220332312030 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2112110110220222-0102223320130213-3033312332110031-3231100202031303-1203102122123002-3303101020302100-3321230301333220-1131300300321312"></a>

## tenant property — segment / 220332312030 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0003223220302120-3232203320333031-2000112302121211-1301103011013303-2010301033022132-3221013311301320-2322313230000130-1302211131222131"></a>

## Next pages — segment / 220332312030 / 7

- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2110032230130213-1233131020100213-0103031331322233-3132300222023320-1221130110230312-3122233023110122-0213032231121113-3000120212301203"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0102322022211200-1033211210333212-1301021012330101-2112210310110111-3233221111001320-0202302331130020-0331332113203201-0022222200122333"></a>

## origin_servers.private_name.site_locator — site_locator / 032223322113 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- origin_servers.private_name.site_locator

<a id="canonical-1323033030323301-1000213232102213-1222113200010031-2233202002201211-0312021122012220-3011323321322103-1222233021120013-3212331312020123"></a>

Type: `"single"`. Computed.

Message defines a reference to a site or virtual site object.

Upstream description:

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

<a id="canonical-3302022310333033-0111021322333203-0103233331230203-3213332303331103-0121232232220133-2020331012300021-0112032000331312-1102001003223020"></a>

## Direct properties — site_locator / 032223322113 / 3

- [site](data-sources--origin_pool--reference--group-002.md#canonical-0032313212303101-0102131210033022-3013302311001130-1212201303121120-0132030031100121-1301121323310103-0023103133011303-1013221113323200): complete subsection reference.

- [virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-0133303123032010-2303131222113200-2231102200221213-2211200102122101-3023311211113031-0011312030031003-1302130011213010-3030330302200322): complete subsection reference.

<a id="canonical-3322032211232233-3312111323021303-2201202331200302-0021223332231320-1121011001123131-0102013101102203-0232021313312300-3021100101323123"></a>

## Next pages — site_locator / 032223322113 / 4

- [origin_servers.private_name.site_locator.site](data-sources--origin_pool--reference--group-002.md#canonical-0032313212303101-0102131210033022-3013302311001130-1212201303121120-0132030031100121-1301121323310103-0023103133011303-1013221113323200)
- [origin_servers.private_name.site_locator.virtual_site](data-sources--origin_pool--reference--group-002.md#canonical-0133303123032010-2303131222113200-2231102200221213-2211200102122101-3023311211113031-0011312030031003-1302130011213010-3030330302200322)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-0032313212303101-0102131210033022-3013302311001130-1212201303121120-0132030031100121-1301121323310103-0023103133011303-1013221113323200"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0303131330212110-0311233033100022-3132002020000232-0122023221122003-3112021332000321-3301332020010103-0203103300133301-2202013112111311"></a>

## origin_servers.private_name.site_locator.site — site / 030011210013 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- [origin_servers.private_name.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-2110032230130213-1233131020100213-0103031331322233-3132300222023320-1221130110230312-3122233023110122-0213032231121113-3000120212301203)
- origin_servers.private_name.site_locator.site

<a id="canonical-2122103123032102-2131333320013322-0213003013010301-3310201212001002-3211000232030113-1101112032102111-0331323101012220-3003001010121201"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-2200203231331210-2032002320223211-2201023031313130-1130232221331302-1203200300201023-3300230220002312-1001003212112330-0101112102321102"></a>

## Direct properties — site / 030011210013 / 3

<a id="canonical-1211233032113220-1001202021313232-0311202130333311-2311300302311102-2231311323110320-0001210120300302-1013101012110102-3130110300020010"></a>

<a id="canonical-0131020200201311-0030300021011221-0303122232320102-0101111302223020-2131121023233331-0232330221031312-2313011321300132-3133130301123111"></a>

## name property — site / 030011210013 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1013212203012333-2313031222323123-3003303222330000-2103010030201132-1011000001202312-0020131122002123-2101023313210121-1103102013201221"></a>

## namespace property — site / 030011210013 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0221223131123132-0130111130203301-3020011000211013-3020032212010203-3212101122323210-1213001121323002-2303320312013122-1323100311313133"></a>

## tenant property — site / 030011210013 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0021321311102320-2210220223023132-0303002103331210-0100103120003330-3032321021232133-3132012321230103-3323100303322133-1303123020333112"></a>

## Next pages — site / 030011210013 / 7

- [origin_servers.private_name.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-2110032230130213-1233131020100213-0103031331322233-3132300222023320-1221130110230312-3122233023110122-0213032231121113-3000120212301203)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-0133303123032010-2303131222113200-2231102200221213-2211200102122101-3023311211113031-0011312030031003-1302130011213010-3030330302200322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0130133110323230-1023312321230323-2021021322011030-3020012001122123-2101112220301323-3312031023103023-3331120030332222-1021233303113002"></a>

## origin_servers.private_name.site_locator.virtual_site — virtual_site / 130221232303 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- [origin_servers.private_name.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-2110032230130213-1233131020100213-0103031331322233-3132300222023320-1221130110230312-3122233023110122-0213032231121113-3000120212301203)
- origin_servers.private_name.site_locator.virtual_site

<a id="canonical-1030132131011010-1212123123013211-0221012011230212-1223121332332322-1032232202022121-3011331310213102-0103033321011323-0012301013230133"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-0330321331222110-0231313120110032-1030211030232202-1132331301300303-2112330332220322-0202311120001311-3112213233031103-2131213322010103"></a>

## Direct properties — virtual_site / 130221232303 / 3

<a id="canonical-3110033030101323-3233302100100230-0200301100201200-0131232232012032-2230100221120101-3122323311002221-0222300103032322-0222100011303122"></a>

<a id="canonical-1021020023123132-1120320000223033-1012301102030330-2021013331102200-3213232323122200-0321023330230312-3210030111210202-3213131132013103"></a>

## name property — virtual_site / 130221232303 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3222021031210022-2311101210133232-3033030010111102-0033330031011212-1201231310233322-2130233301023002-3013303101300031-3021123011202123"></a>

## namespace property — virtual_site / 130221232303 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3221203033230003-2012001322121232-1020311332322203-2012330000230303-0130020023032013-2021122211022102-1331322112033123-3201210110303201"></a>

## tenant property — virtual_site / 130221232303 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3113203330003112-1110200002113130-1232321100023303-2333123133300113-0022230131330101-0300222233230103-1131203123011320-1303112232001000"></a>

## Next pages — virtual_site / 130221232303 / 7

- [origin_servers.private_name.site_locator](data-sources--origin_pool--reference--group-002.md#canonical-2110032230130213-1233131020100213-0103031331322233-3132300222023320-1221130110230312-3122233023110122-0213032231121113-3000120212301203)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-3312011332131233-1230133131313010-3013003121231323-0223212213120012-0301213303101202-1000312320032330-3210122031212113-2000030222032011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1001001321300102-0012323120021223-1221133112013130-0311003232311112-0311313101023312-0310213022300313-2031223332122010-3221301330321010"></a>

## origin_servers.private_name.snat_pool — snat_pool / 221303212211 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- origin_servers.private_name.snat_pool

<a id="canonical-1331233212211320-2330123312023213-0133103211121212-0120101132022101-0030332333002120-1303321100121230-3122021023133032-1301312121300213"></a>

Type: `"single"`. Computed.

SNAT Pool. SNAT Pool configuration.

Upstream description:

SNAT Pool configuration.

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

<a id="canonical-3311322133303332-3031022133013221-1300123230123033-1323212022313111-0102200310101132-0300103321322303-1213013212311021-1001131213230233"></a>

## Direct properties — snat_pool / 221303212211 / 3

- [no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-3312120121220023-3230131203113023-1222013132133331-2010313033233230-2332333210322111-2102020101101330-2111331013210333-1133201121322223): complete subsection reference.

- [snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-3212302120120200-1323200000301103-0223232010213033-3103221203300131-0131233211123212-0100110203310302-1021032313011123-3101121333122003): complete subsection reference.

<a id="canonical-2000000331220033-3233031221311200-2301112013312331-0000130223320101-0122313202011030-2020120213223303-1322012323022020-0000111320131120"></a>

## Next pages — snat_pool / 221303212211 / 4

- [origin_servers.private_name.snat_pool.no_snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-3312120121220023-3230131203113023-1222013132133331-2010313033233230-2332333210322111-2102020101101330-2111331013210333-1133201121322223)
- [origin_servers.private_name.snat_pool.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-3212302120120200-1323200000301103-0223232010213033-3103221203300131-0131233211123212-0100110203310302-1021032313011123-3101121333122003)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-3312120121220023-3230131203113023-1222013132133331-2010313033233230-2332333210322111-2102020101101330-2111331013210333-1133201121322223"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3321232200111221-0300210000332000-3123223220311103-1332033033333232-2100203322330213-0121022201332311-2033223230302203-3333010000000123"></a>

## origin_servers.private_name.snat_pool.no_snat_pool — no_snat_pool / 032331000101 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- [origin_servers.private_name.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-3312011332131233-1230133131313010-3013003121231323-0223212213120012-0301213303101202-1000312320032330-3210122031212113-2000030222032011)
- origin_servers.private_name.snat_pool.no_snat_pool

<a id="canonical-1231201333130132-1002102100110300-1031113030320232-0200000330233032-0012221031123231-3120101323313200-3122130021321332-3203010113331223"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for no snat pool.

Upstream description:

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

<a id="canonical-1300221120313223-2221220233223310-2101020010201321-1333332022013332-1330211122132001-0302301332203112-2120212131121321-3002030013122201"></a>

## Direct properties — no_snat_pool / 032331000101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2001020130303221-2000311031323200-3103132310330323-3131131301032320-0321332322223310-3202330300001313-0120103312302012-1302002021331100"></a>

## Next pages — no_snat_pool / 032331000101 / 4

- [origin_servers.private_name.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-3312011332131233-1230133131313010-3013003121231323-0223212213120012-0301213303101202-1000312320032330-3210122031212113-2000030222032011)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-3212302120120200-1323200000301103-0223232010213033-3103221203300131-0131233211123212-0100110203310302-1021032313011123-3101121333122003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3021110312231313-1013223300320203-2123231113002023-2221132132211320-1013122322203302-2330011123223111-1333032131302103-3301301121030221"></a>

## origin_servers.private_name.snat_pool.snat_pool — snat_pool / 221202331021 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.private_name](data-sources--origin_pool--reference--group-002.md#canonical-1201220111301320-0211203213002232-0010033301302122-1330201020221332-2020323033101033-2023122011200330-2102012001330301-3002003012033110)
- [origin_servers.private_name.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-3312011332131233-1230133131313010-3013003121231323-0223212213120012-0301213303101202-1000312320032330-3210122031212113-2000030222032011)
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

<a id="canonical-1222132033333330-3100101032100002-3313202123331323-1110003220302212-0300132023321131-1333223233221201-0012310323122203-1322311001330202"></a>

## Direct properties — snat_pool / 221202331021 / 3

<a id="canonical-0233103211231133-2110222003232101-0320133301200002-3030012213211320-1033011002033131-3300111023323230-1211310033212312-3120300302301312"></a>

<a id="canonical-0231220310021101-1223000232122233-3222021010002323-2311110102010102-2033113133200301-1303213033320032-0313233300221330-3330321012111222"></a>

## prefixes property — snat_pool / 221202331021 / 4

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1020111201011222-1113010030100103-3022333121002000-2220303310023002-0331231111001220-3302100103113130-3031202313322331-0122222111213031"></a>

## Next pages — snat_pool / 221202331021 / 5

- [origin_servers.private_name.snat_pool](data-sources--origin_pool--reference--group-002.md#canonical-3312011332131233-1230133131313010-3013003121231323-0223212213120012-0301213303101202-1000312320032330-3210122031212113-2000030222032011)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-3233223332020121-3100021120202320-2031230003101211-3023100123103102-0220333130210323-2201111232211122-1033211313103113-3001033210113321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2323110231011302-2133120002000113-3310121003110102-2301022313300322-2300033101311230-3021111220132103-3022120223131333-0031031230200000"></a>

## origin_servers.public_ip — public_ip / 312320213022 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
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

<a id="canonical-1131033021121330-1333133323212021-3103202110213220-1130233102331211-3320122100001012-2220011010101112-2313133013002011-0002111013131332"></a>

## Direct properties — public_ip / 312320213022 / 3

<a id="canonical-2112313312012323-2310022330233310-3112311113003102-1303230123010002-1331011111322201-0333202020101120-3201102313122030-0100112212232122"></a>

<a id="canonical-3310312101103330-1222223010220210-3131122013302121-1112011133223232-2201121230331233-2130231311122010-0303332310101212-1320110010200203"></a>

## ip property — public_ip / 312320213022 / 4

Type: `"string"`. Computed.

Public IPv4. Exclusive with \[\] Public IPv4 address.

Upstream description:

Exclusive with \[\] Public IPv4 address.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3322033233010323-2133332000031323-1232101013110123-3122121231312102-3200231001100221-0113212033303201-0100023222020003-0030222233213011"></a>

## Next pages — public_ip / 312320213022 / 5

- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-3023121012300320-1113333033031232-1003130321112112-1203102212132000-1310200323133230-2031320320302202-3121312332032031-2323130012131111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1323210302302033-2302210210110032-0013213302210011-2131220123311320-2213232031123110-2001132112122322-0200231210000300-2231110330123220"></a>

## origin_servers.public_name — public_name / 220012130203 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
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

<a id="canonical-2012031011312012-2231302022303020-1233021222320211-3332303111222200-0213023101002030-3000002112303301-1300210033230232-2133331211030032"></a>

## Direct properties — public_name / 220012130203 / 3

<a id="canonical-1200212100313122-0230222122110333-3032013211322133-2223123001330111-1103312130303231-1030200121030133-0301323223203322-3302113023330332"></a>

<a id="canonical-3132303120300111-2020331320101201-2100012210210212-0001223011022210-2212320120031001-0120002002301102-1323302332232032-2111212300003001"></a>

## dns_name property — public_name / 220012130203 / 4

Type: `"string"`. Computed.

DNS Name. DNS Name

Upstream description:

DNS Name

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0002332123100211-0211320323013022-3221103212220222-3100311033311321-3323222102200011-0303321113100222-2230221323021211-1210330033333211"></a>

## refresh_interval property — public_name / 220012130203 / 5

Type: `"number"`. Computed.

Interval for DNS refresh in seconds. Max value is 7 days as per
https&#58;//datatracker.ietf.org/doc/HTML/rfc8767.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-0122301021123233-1233320232300112-2031031122303223-2210010201203230-3210123303301233-0013222032103011-2031203332333102-1220313320013213"></a>

## Next pages — public_name / 220012130203 / 6

- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-1132312333312222-3332122302332200-2310202220131303-2312323300201122-2030322023323212-2010332131123003-0331231022031330-0313233113021202"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1213310301123132-0312133022211301-3213021120200000-2001333322300100-1122023321210123-2330311033031310-0102130022132003-3020131033230312"></a>

## origin_servers.vn_private_ip — vn_private_ip / 011302123320 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
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

<a id="canonical-0330221123311212-2223023111311232-1132031100121322-2212213221332230-0233000020323302-2201101113012310-0102230113202022-1032230211312102"></a>

## Direct properties — vn_private_ip / 011302123320 / 3

<a id="canonical-2003223332223001-3232311123032310-1312001010002300-2030002132223000-0202011113321011-3220300313123232-2312120221301013-2021130320222030"></a>

<a id="canonical-1022310132101030-2033310301121123-0102030110221331-3222123331013333-1000101130310011-0332021311222111-2013001120130112-2123013233000303"></a>

## ip property — vn_private_ip / 011302123320 / 4

Type: `"string"`. Computed.

IPv4. Exclusive with \[\] IPv4 address.

Upstream description:

Exclusive with \[\] IPv4 address.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [virtual_network](data-sources--origin_pool--reference--group-002.md#canonical-2010102112110023-3022011022122033-2323023203111013-3233300030300132-3122331230123120-1200000330211011-0221022301313020-2112231321100023): complete subsection reference.

<a id="canonical-3301110321303222-1120322112201103-0232301100211332-2013233321030333-3102111100021322-3202232003220032-3021310332132213-3310232121221332"></a>

## Next pages — vn_private_ip / 011302123320 / 5

- [origin_servers.vn_private_ip.virtual_network](data-sources--origin_pool--reference--group-002.md#canonical-2010102112110023-3022011022122033-2323023203111013-3233300030300132-3122331230123120-1200000330211011-0221022301313020-2112231321100023)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2010102112110023-3022011022122033-2323023203111013-3233300030300132-3122331230123120-1200000330211011-0221022301313020-2112231321100023"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010301311333200-1203100322021302-2011300301330010-0033303122003213-1333031222210010-0132031233210300-2100101022122310-2022001230222222"></a>

## origin_servers.vn_private_ip.virtual_network — virtual_network / 120032200313 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.vn_private_ip](data-sources--origin_pool--reference--group-002.md#canonical-1132312333312222-3332122302332200-2310202220131303-2312323300201122-2030322023323212-2010332131123003-0331231022031330-0313233113021202)
- origin_servers.vn_private_ip.virtual_network

<a id="canonical-0002021132223321-2003211320303002-1103121321312311-2223210300112331-0300321112200111-3133303122132323-0323330102103021-3030211012303332"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-1221210201032100-3003001223120103-0313232201310310-0122101103103101-0200010130332022-3002022101210100-1310012133311202-1220302211000132"></a>

## Direct properties — virtual_network / 120032200313 / 3

<a id="canonical-1012013330302213-0231110132123330-2133213131023301-2023331220323003-3232332330210031-2130022132022322-2300311020013000-3203103032110201"></a>

<a id="canonical-1133031310131112-2000233120221213-0133022200002131-2323112311023233-1201222130211212-2112121310123301-3011200332110120-3230100012010321"></a>

## name property — virtual_network / 120032200313 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3103303032122333-0110230002003110-0222330120023102-2111311311031232-2011202221200311-1123222110312101-3312120320122102-3320231013201230"></a>

## namespace property — virtual_network / 120032200313 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1021111333033121-1100023011002222-0210202320210303-3222130212203011-3313202332301011-3023330133232223-2022113130321131-0210022022033000"></a>

## tenant property — virtual_network / 120032200313 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3320020021311233-0132232101120321-3231102210023213-0021310230033221-3030011311223321-3222013021021111-2132200333200033-3103212211310013"></a>

## Next pages — virtual_network / 120032200313 / 7

- [origin_servers.vn_private_ip](data-sources--origin_pool--reference--group-002.md#canonical-1132312333312222-3332122302332200-2310202220131303-2312323300201122-2030322023323212-2010332131123003-0331231022031330-0313233113021202)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-3322112003203211-1300331030110100-3211113003210110-3003332110030121-2200102131333003-3323033022212213-0033103302220322-0021203022320330"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0001111201100100-1323321112232332-0212212102333200-3321003121013220-0203113331001132-0020231301201023-3133023013312201-0011021032303132"></a>

## origin_servers.vn_private_name — vn_private_name / 031120020312 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
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

<a id="canonical-1023000132023332-1330123311131033-2020023332122331-2311303333203131-0223030201222103-1102310101020100-0011013332101321-1121102100231020"></a>

## Direct properties — vn_private_name / 031120020312 / 3

<a id="canonical-1101311331303030-0032012310223201-1030113001021100-3223200320231312-2300221231321110-3211030022002012-0323333220300300-1120312102213313"></a>

<a id="canonical-3300311012332022-1103211000321321-2012230213320022-2010330300201101-1322133231110323-2010233130100110-3131212012013113-2333021010023313"></a>

## dns_name property — vn_private_name / 031120020312 / 4

Type: `"string"`. Computed.

DNS Name. DNS Name

Upstream description:

DNS Name

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

- [private_network](data-sources--origin_pool--reference--group-002.md#canonical-2022222020213202-3110010323312231-2200233032000223-2213123322022120-1010112202031033-1313212110330102-3021310121230201-1103102210232313): complete subsection reference.

<a id="canonical-0120220333022123-0011012322102333-1203001030103033-3330302231300021-1032113132333100-2233211103123110-3031311323023201-0331330313112113"></a>

## Next pages — vn_private_name / 031120020312 / 5

- [origin_servers.vn_private_name.private_network](data-sources--origin_pool--reference--group-002.md#canonical-2022222020213202-3110010323312231-2200233032000223-2213123322022120-1010112202031033-1313212110330102-3021310121230201-1103102210232313)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2022222020213202-3110010323312231-2200233032000223-2213123322022120-1010112202031033-1313212110330102-3021310121230201-1103102210232313"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2233012333322333-2310301302113132-1303331012001212-2132121323233233-0221230320122021-0101212120032231-3031113200311300-2000313003233023"></a>

## origin_servers.vn_private_name.private_network — private_network / 323102113322 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [origin_servers](data-sources--origin_pool--reference--group-001.md#canonical-1233300011221103-2310303122111113-0230213330333210-0333121303222200-3102122111111200-0213023003023223-0021231311232231-1123030010132211)
- [origin_servers.vn_private_name](data-sources--origin_pool--reference--group-002.md#canonical-3322112003203211-1300331030110100-3211113003210110-3003332110030121-2200102131333003-3323033022212213-0033103302220322-0021203022320330)
- origin_servers.vn_private_name.private_network

<a id="canonical-2022001023201223-1213310013210002-2333312233011020-1133010220111313-0232120012022232-2322002013233023-3321233023222332-3030101333030010"></a>

Type: `"single"`. Computed.

Type establishes a direct reference from one object(the referrer) to another(the referred). Such a
reference is in form of tenant/namespace/name.

Upstream description:

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

<a id="canonical-2310212220231003-1011323221110330-1032212320033322-0001322301322012-3020122023003200-1120333122231220-0110022330232033-0222133220001210"></a>

## Direct properties — private_network / 323102113322 / 3

<a id="canonical-3131030122030032-2223021030311310-3020011103310230-0100010132233300-3332032022303322-0313220320031220-3311102032113020-2333232131200011"></a>

<a id="canonical-0012003213112201-3102001030111132-0000021131203331-2221323113301112-3112000310020202-3331311212211032-0211221331013121-0213232031230013"></a>

## name property — private_network / 323102113322 / 4

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then name will hold the
referred object's(e.g. Route's) name.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1333221332333110-2001133130330010-1103312030002031-0111010313023002-0022230102010222-0131312133101203-3203023033000000-3113013101111120"></a>

## namespace property — private_network / 323102113322 / 5

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then namespace will
hold the referred object's(e.g. Route's) namespace.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-2302223312332221-1000111331202031-3123300211010212-0120221223313310-1330201230213330-2101323302003223-1103022010202202-3233011232200222"></a>

## tenant property — private_network / 323102113322 / 6

Type: `"string"`. Computed.

When a configuration object(e.g. Virtual\_host) refers to another(e.g route) then tenant will hold
the referred object's(e.g. Route's) tenant.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-1330332220333211-2313131233231220-0223311002322223-0311303112002031-0303100030203333-1222321331012332-3223001310010033-3312013132222023"></a>

## Next pages — private_network / 323102113322 / 7

- [origin_servers.vn_private_name](data-sources--origin_pool--reference--group-002.md#canonical-3322112003203211-1300331030110100-3211113003210110-3003332110030121-2200102131333003-3323033022212213-0033103302220322-0021203022320330)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-3130012011302310-0100000020300011-2333213101123221-2323300231131231-3230213133122123-3221221202130031-0122330100113130-1102320312113033"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2010111200033103-1320132230113102-0323032101103031-0002022223201103-1001120001323211-3222221102000332-1010031121303232-2012032301202020"></a>

## same_as_endpoint_port — same_as_endpoint_port / 103022201003 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- same_as_endpoint_port

<a id="canonical-3230202001231011-0231003131211320-3300230323330013-3310330210031113-0323111321313021-1120103003020320-2233213133003321-0013202003001030"></a>

Type: `["object", {}]`. Computed.

Enable this option. Defaults to \`map\[\]\`. Server applies default when omitted.

Upstream description:

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

<a id="canonical-0001230223021200-0301211302230010-1301323031020301-2323130320110201-1002310313211130-2111230011123103-0020031001213333-2103020323102112"></a>

## Direct properties — same_as_endpoint_port / 103022201003 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2002303202230310-1122321323000221-3133300113011322-0123223103233321-3330202131033111-2021102203100030-3012223132000101-3011012330211210"></a>

## Next pages — same_as_endpoint_port / 103022201003 / 4

- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-2231230203233120-1223102033221212-2211002030000333-1332212222223122-0031021313330101-0202130201022222-3131132023113300-0330011232102031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1222110232330200-1312110123003130-0301333113011310-2132122023323032-0211322323023312-2010232203203231-0022023311302022-2230322022101133"></a>

## upstream_conn_pool_reuse_type — upstream_conn_pool_reuse_type / 010132101302 / 2

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

<a id="canonical-1201032031011230-0121123000123202-0133303031221112-0131110302200011-3013232233111020-3301032022003332-0101320203202210-3011330110132323"></a>

## Direct properties — upstream_conn_pool_reuse_type / 010132101302 / 3

- [disable_conn_pool_reuse](data-sources--origin_pool--reference--group-002.md#canonical-1022200321320132-1123023332011133-2220201012202003-2200133202320303-0210110310113332-3002003313012203-0131002220211110-2221203011111120): complete subsection reference.

- [enable_conn_pool_reuse](data-sources--origin_pool--reference--group-002.md#canonical-1320101022032011-3203201130311130-0232121211022033-3203300130321003-3323120121012121-0112332312130333-0000133030030300-2012001123130123): complete subsection reference.

<a id="canonical-2101221120232323-2122130021031011-2213221213221102-0122033223031130-2012212000322033-3323011322002233-0030112023232022-3103101203100300"></a>

## Next pages — upstream_conn_pool_reuse_type / 010132101302 / 4

- [upstream_conn_pool_reuse_type.disable_conn_pool_reuse](data-sources--origin_pool--reference--group-002.md#canonical-1022200321320132-1123023332011133-2220201012202003-2200133202320303-0210110310113332-3002003313012203-0131002220211110-2221203011111120)
- [upstream_conn_pool_reuse_type.enable_conn_pool_reuse](data-sources--origin_pool--reference--group-002.md#canonical-1320101022032011-3203201130311130-0232121211022033-3203300130321003-3323120121012121-0112332312130333-0000133030030300-2012001123130123)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-1022200321320132-1123023332011133-2220201012202003-2200133202320303-0210110310113332-3002003313012203-0131002220211110-2221203011111120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0003022203000030-1322110130013013-3032121331203131-1020102213101013-2230031203032320-2102112000012203-0000102000322223-3200322221202021"></a>

## upstream_conn_pool_reuse_type.disable_conn_pool_reuse — disable_conn_pool_reuse / 003221300332 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [upstream_conn_pool_reuse_type](data-sources--origin_pool--reference--group-002.md#canonical-2231230203233120-1223102033221212-2211002030000333-1332212222223122-0031021313330101-0202130201022222-3131132023113300-0330011232102031)
- upstream_conn_pool_reuse_type.disable_conn_pool_reuse

<a id="canonical-1221113302121301-2313020121331320-0221203333123323-2123200223300002-1330031110322133-3313200133100312-1310111322321002-2123221332010213"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable conn pool reuse.

Upstream description:

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

<a id="canonical-3222032221002202-1101023030321201-3120221113331110-2332133233210022-1333312122311120-1232102121333321-1012122312233200-1212202200320112"></a>

## Direct properties — disable_conn_pool_reuse / 003221300332 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1013312021013033-0011110120221322-2230031210223300-3220220003002313-1221211200003301-2020122323202202-3103011301123212-1123000330122210"></a>

## Next pages — disable_conn_pool_reuse / 003221300332 / 4

- [upstream_conn_pool_reuse_type](data-sources--origin_pool--reference--group-002.md#canonical-2231230203233120-1223102033221212-2211002030000333-1332212222223122-0031021313330101-0202130201022222-3131132023113300-0330011232102031)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-1320101022032011-3203201130311130-0232121211022033-3203300130321003-3323120121012121-0112332312130333-0000133030030300-2012001123130123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300321232311201-1100201121132012-2133022202200013-2031111012333221-3331123103301013-2331202222132323-0231101222303123-0020233333102331"></a>

## upstream_conn_pool_reuse_type.enable_conn_pool_reuse — enable_conn_pool_reuse / 130112103101 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [upstream_conn_pool_reuse_type](data-sources--origin_pool--reference--group-002.md#canonical-2231230203233120-1223102033221212-2211002030000333-1332212222223122-0031021313330101-0202130201022222-3131132023113300-0330011232102031)
- upstream_conn_pool_reuse_type.enable_conn_pool_reuse

<a id="canonical-1303312032131320-2021212201121011-2021212032121302-0213111331132120-0201200332100220-0201322223303113-2130120213131233-3303120103001132"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for enable conn pool reuse.

Upstream description:

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

<a id="canonical-3303222030100023-1113323202323111-3020320113331220-3003133230321003-2102310031313123-2211123003030101-1211301010221002-1122033133113321"></a>

## Direct properties — enable_conn_pool_reuse / 130112103101 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-2223101030310002-0311310132231130-1231132102320201-1103223303311200-1303322030013031-2001003223220000-0231202230222130-3011302221103130"></a>

## Next pages — enable_conn_pool_reuse / 130112103101 / 4

- [upstream_conn_pool_reuse_type](data-sources--origin_pool--reference--group-002.md#canonical-2231230203233120-1223102033221212-2211002030000333-1332212222223122-0031021313330101-0202130201022222-3131132023113300-0330011232102031)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1220221332111002-3221322023100232-3112031001000133-0231121231332321-1202012120331222-0013133323002211-2131203222110102-2231231000003222"></a>

## use_tls — use_tls / 011213013312 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- use_tls

<a id="canonical-0002222103101132-0331331011030000-0221212201230213-1311030121013132-0223020023313332-2311112031102200-3132032023331103-2122132001320021"></a>

Type: `"single"`. Computed.

TLS Parameters for Origin Servers. Upstream TLS Parameters.

Upstream description:

Upstream TLS Parameters.

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

<a id="canonical-0013302310100031-3211121021101022-0100131102103311-2321131313002001-1102232200012121-0233301101233130-3323210312130232-2213303011030310"></a>

## Direct properties — use_tls / 011213013312 / 3

- [default_session_key_caching](data-sources--origin_pool--reference--group-002.md#canonical-3120033223300122-0002012021322010-3133211003121300-2201320200112013-1022121320122111-1321231300011022-1202030223010131-2130002001033012): complete subsection reference.

- [disable_session_key_caching](data-sources--origin_pool--reference--group-003.md#canonical-3112023032132312-1132010002311002-3003320230110132-1123303202301332-0312123102222332-0303130310001332-1123133313201330-2001030020320211): complete subsection reference.

- [disable_sni](data-sources--origin_pool--reference--group-003.md#canonical-3302331001011103-0032210313332130-2330101100323023-3233031132332300-3233222030222021-0231223303003021-1220133302031312-1211313003212132): complete subsection reference.

<a id="canonical-0233103101220320-1020311303110113-1100311111211321-1231332102301202-2320032130222232-0033111113112220-2100121012331321-0130112230133010"></a>

<a id="canonical-0123001130223101-0321300122302323-2200010121121213-1131021231333302-0020131300000322-3331112312122331-2100223301303113-3312310131312200"></a>

## max_session_keys property — use_tls / 011213013312 / 4

Type: `"number"`. Computed.

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\] Number of session
keys that are cached.

Upstream description:

Exclusive with \[default\_session\_key\_caching disable\_session\_key\_caching\]

Number of session keys that are cached.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3130211001012223-3013232101313013-0301020011230311-0000133031023233-1011112030233013-0220231020130320-0002303230302233-2300321323310102"></a>

## sni property — use_tls / 011213013312 / 5

Type: `"string"`. Computed.

Exclusive with \[disable\_sni use\_host\_header\_as\_sni\] SNI value to be used.

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

<a id="canonical-3212021003023320-3131010012032231-1323130120113203-0311231320133120-1230310313210011-1233211001232323-0011001210302230-2101330202021311"></a>

## Next pages — use_tls / 011213013312 / 6

- [use_tls.default_session_key_caching](data-sources--origin_pool--reference--group-002.md#canonical-3120033223300122-0002012021322010-3133211003121300-2201320200112013-1022121320122111-1321231300011022-1202030223010131-2130002001033012)
- [use_tls.disable_session_key_caching](data-sources--origin_pool--reference--group-003.md#canonical-3112023032132312-1132010002311002-3003320230110132-1123303202301332-0312123102222332-0303130310001332-1123133313201330-2001030020320211)
- [use_tls.disable_sni](data-sources--origin_pool--reference--group-003.md#canonical-3302331001011103-0032210313332130-2330101100323023-3233031132332300-3233222030222021-0231223303003021-1220133302031312-1211313003212132)
- [use_tls.no_mtls](data-sources--origin_pool--reference--group-003.md#canonical-2033030130120230-2002013012023100-2300101011010123-0023110121222313-2002022101100132-1200212102011022-0232110202003010-1021021323033132)
- [use_tls.skip_server_verification](data-sources--origin_pool--reference--group-003.md#canonical-3133102211002303-3032200122321120-0210020232102320-2021111001013011-1213213100231100-2122021233030322-3321111131233310-0313010303321021)
- [use_tls.tls_config](data-sources--origin_pool--reference--group-003.md#canonical-3002221101212230-2032212202012222-2211102302333323-3031111211101332-0202202331023033-3323013021210110-1130132200330201-1122230100213213)
- [use_tls.use_host_header_as_sni](data-sources--origin_pool--reference--group-003.md#canonical-3232122223122133-3321210122301001-0222031200332131-2203033233132002-0222102221311230-2103030133101313-3220313220123111-1321233321311021)
- [use_tls.use_mtls](data-sources--origin_pool--reference--group-003.md#canonical-2123230122313111-0311121021121133-1321002130011032-3312300200211002-2110033332211203-2121202220332232-0100223311202100-1321101100003310)
- [use_tls.use_mtls_obj](data-sources--origin_pool--reference--group-003.md#canonical-3310130231121301-1123130000302011-2213030223021110-2221211022020003-3033112020011311-2022011333233203-1120313312223010-3003101030232320)
- [use_tls.use_server_verification](data-sources--origin_pool--reference--group-003.md#canonical-3013231000221313-0112303320203120-0231130102303320-1100002233231301-1222203202030202-1113323222002130-0331312102131100-0102330131222110)
- [use_tls.volterra_trusted_ca](data-sources--origin_pool--reference--group-003.md#canonical-1332301321201201-2112102322122113-3211233001322320-1230113120130000-0130212221102210-3323230123233332-1121303303011030-0131110111202333)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)

<a id="canonical-3120033223300122-0002012021322010-3133211003121300-2201320200112013-1022121320122111-1321231300011022-1202030223010131-2130002001033012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0333032101331210-3333132000022130-3333033030222032-2233000321001211-0101212122311010-2010223222103221-0231002230323300-3321021200212102"></a>

## use_tls.default_session_key_caching — default_session_key_caching / 032323033303 / 2

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
- [Property reference](data-sources--origin_pool--reference--group-001.md#canonical-0103220113112230-2012100232223113-1023222131101333-0330000000001221-3000302121221221-1110013322303122-0231312123220320-3022030100223303)
- [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- use_tls.default_session_key_caching

<a id="canonical-2322202020222130-1232110023322133-3333101113001113-3210232200231331-1231120123332320-3010301033212312-3103012211310233-0311213210031323"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default session key caching. Defaults to \`map\[\]\`. Server applies
default when omitted.

Upstream description:

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

<a id="canonical-3020032002202330-0331123010103221-0203020123220111-3322113120101013-1231302302220213-2130121132102110-3002132200303000-1202103133033200"></a>

## Direct properties — default_session_key_caching / 032323033303 / 3

This is an empty object or choice marker. It has no direct properties.

<a id="canonical-1133222113003011-2020213033303003-2210333310030010-3310002100101213-1001020103221301-2020233300210300-0123223220123231-3001330122013003"></a>

## Next pages — default_session_key_caching / 032323033303 / 4

- [use_tls](data-sources--origin_pool--reference--group-002.md#canonical-3220233020111000-3232323231033132-3022110210002213-3223110300110302-2322121221301313-3002213311302131-1113131021311021-3120313011133301)
- [xcsh_origin_pool](../data-sources/origin_pool.md#canonical-3103011022300321-3200033030203023-3322013132331030-0120200333302110-1012310120003032-1223333313230200-2333321302032332-1311021030332203)
