const BREAKING_KEYWORDS = ["BREAKING CHANGE", "BREAKING CHANGES", "BREAKING"];
module.exports = {
  branches: [
    "main",
    {
      name: "next",
      prerelease: true,
    },
    {
      name: "alpha",
      prerelease: true,
    },
  ],
  plugins: [
    [
      "@semantic-release/commit-analyzer",
      {
        preset: "angular",
        parserOpts: {
          noteKeywords: BREAKING_KEYWORDS,
        },
      },
    ],
    [
      "@semantic-release/release-notes-generator",
      {
        preset: "angular",
        parserOpts: {
          noteKeywords: BREAKING_KEYWORDS,
        },
      },
    ],
    [
      "@semantic-release/github",
      {
        // assets: [
        //   {
        //     path: "javadoc.zip",
        //     label: "javadoc folder added to release",
        //   },
        //   {
        //     path: "build/libs/*.war",
        //     label: "war app added to release",
        //   },
        // ],
      },
    ],
  ],
};
