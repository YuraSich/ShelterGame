const basePresets = [
    [
        "@babel/preset-env",
        {
            "modules": false
        }
    ]
];

const basePlugins = [
    "@babel/plugin-syntax-dynamic-import",
    "@babel/plugin-syntax-import-meta",
    "@babel/plugin-proposal-class-properties",
    "@babel/plugin-proposal-json-strings",
];

module.exports = function (api) {
    api.cache(true); // necessary

    if (process.env.NODE_ENV === 'development') {

    } else if (process.env.NODE_ENV === 'production') {
        basePlugins.push("transform-remove-console");   // Удаляем onsole.log
    }
    return {
        presets: basePresets,
        plugins: basePlugins,
    }
};
//
// {
//   "presets": [
//   [
//     "@babel/preset-env",
//     {
//       "modules": false
//     }
//   ]
// ],
//     "plugins": [
//   "@babel/plugin-syntax-dynamic-import",
//   "@babel/plugin-syntax-import-meta",
//   "@babel/plugin-proposal-class-properties",
//   "@babel/plugin-proposal-json-strings",
//   "transform-remove-console"
// ]
// }