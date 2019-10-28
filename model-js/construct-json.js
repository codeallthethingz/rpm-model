exports.maxArchiveSizeBytes = 1000000
exports.modelVersion = '0.2.1'
exports.predefinedUnits = [
  'mm',
  'cm',
  'm',
  'km',
  'inch',
  'foot',
  'yard',
  'mile'
]

exports.predefinedBoundingTypes = [
  {"name":"triangular-prism","measurements":{"height":"","radius":""}},
  {"name":"cuboid","measurements":{"height":"","length":"","width":""}},
  {"name":"pentagonal-prism","measurements":{"height":"","radius":""}},
  {"name":"hexagonal-prism","measurements":{"height":"","radius":""}},
  {"name":"heptagonal-prism","measurements":{"height":"","radius":""}},
  {"name":"octagonal-prism","measurements":{"height":"","radius":""}},
  {"name":"nonagon-prism","measurements":{"height":"","radius":""}},
  {"name":"pecagonal-prism","measurements":{"height":"","radius":""}},
  {"name":"round-cylinder","measurements":{"height":"","radius":""}},
  {"name":"oval-cylinder","measurements":{"height":"","radius1":"","radius2":""}},
  {"name":"sphere","measurements":{"radius":""}},
  {"name":"ellipsoid","measurements":{"radius1":"","radius2":""}}
]
