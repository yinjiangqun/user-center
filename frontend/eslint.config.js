import eslint from "@eslint/js";
import tseslint from "@typescript-eslint/eslint-plugin";
import * as parser from "@typescript-eslint/parser";
import * as vueParser from "vue-eslint-parser";

export default [
  eslint.configs.recommended,
  {
    files: ["**/*.{ts,tsx,vue}"],
    languageOptions: {
      parser: vueParser,
      parserOptions: {
        parser: parser,
        ecmaVersion: 2020,
        sourceType: "module",
        extraFileExtensions: [".vue"],
      },
    },
    plugins: {
      "@typescript-eslint": tseslint,
    },
    rules: {
      "no-console": process.env.NODE_ENV === "production" ? "warn" : "off",
      "no-debugger": process.env.NODE_ENV === "production" ? "warn" : "off",
    },
  },
];
