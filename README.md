# Golang Pagespeed API Example

To get an API key visit [PageSpeed Insights API](https://developers.google.com/speed/docs/insights/v5/get-started) page

Then you need to change **.env.example** file.

That's all :)

An example result here:

```json
{
  "captchaResult": "CAPTCHA_NOT_NEEDED",
  "kind": "pagespeedonline#result",
  "id": "https://aligoren.com/",
  "loadingExperience": {
    "id": "https://aligoren.com",
    "metrics": {
      "CUMULATIVE_LAYOUT_SHIFT_SCORE": {
        "percentile": 0,
        "distributions": [
          {
            "min": 0,
            "max": 10,
            "proportion": 0.9602958937198067
          },
          {
            "min": 10,
            "max": 25,
            "proportion": 0.008907004830917874
          },
          {
            "min": 25,
            "max": 0,
            "proportion": 0.030797101449275378
          }
        ],
        "category": "FAST"
      },
      "EXPERIMENTAL_INTERACTION_TO_NEXT_PAINT": {
        "percentile": 35,
        "distributions": [
          {
            "min": 0,
            "max": 200,
            "proportion": 0.9615074891286842
          },
          {
            "min": 200,
            "max": 500,
            "proportion": 0.02045417941697536
          },
          {
            "min": 500,
            "max": 0,
            "proportion": 0.01803833145434048
          }
        ],
        "category": "FAST"
      },
      "EXPERIMENTAL_TIME_TO_FIRST_BYTE": {
        "percentile": 619,
        "distributions": [
          {
            "min": 0,
            "max": 800,
            "proportion": 0.7944588917783557
          },
          {
            "min": 800,
            "max": 1800,
            "proportion": 0.1080216043208642
          },
          {
            "min": 1800,
            "max": 0,
            "proportion": 0.09751950390077999
          }
        ],
        "category": "FAST"
      },
      "FIRST_CONTENTFUL_PAINT_MS": {
        "percentile": 1521,
        "distributions": [
          {
            "min": 0,
            "max": 1800,
            "proportion": 0.8027352460234876
          },
          {
            "min": 1800,
            "max": 3000,
            "proportion": 0.13349189832020217
          },
          {
            "min": 3000,
            "max": 0,
            "proportion": 0.0637728556563103
          }
        ],
        "category": "FAST"
      },
      "LARGEST_CONTENTFUL_PAINT_MS": {
        "percentile": 1643,
        "distributions": [
          {
            "min": 0,
            "max": 2500,
            "proportion": 0.8943399044205497
          },
          {
            "min": 2500,
            "max": 4000,
            "proportion": 0.08624551971326162
          },
          {
            "min": 4000,
            "max": 0,
            "proportion": 0.019414575866188773
          }
        ],
        "category": "FAST"
      }
    },
    "overall_category": "FAST",
    "initial_url": "https://aligoren.com/",
    "origin_fallback": true
  },
  "originLoadingExperience": {
    "id": "https://aligoren.com",
    "metrics": {
      "CUMULATIVE_LAYOUT_SHIFT_SCORE": {
        "percentile": 0,
        "distributions": [
          {
            "min": 0,
            "max": 10,
            "proportion": 0.9602958937198067
          },
          {
            "min": 10,
            "max": 25,
            "proportion": 0.008907004830917874
          },
          {
            "min": 25,
            "max": 0,
            "proportion": 0.030797101449275378
          }
        ],
        "category": "FAST"
      },
      "EXPERIMENTAL_INTERACTION_TO_NEXT_PAINT": {
        "percentile": 35,
        "distributions": [
          {
            "min": 0,
            "max": 200,
            "proportion": 0.9615074891286842
          },
          {
            "min": 200,
            "max": 500,
            "proportion": 0.02045417941697536
          },
          {
            "min": 500,
            "max": 0,
            "proportion": 0.01803833145434048
          }
        ],
        "category": "FAST"
      },
      "EXPERIMENTAL_TIME_TO_FIRST_BYTE": {
        "percentile": 619,
        "distributions": [
          {
            "min": 0,
            "max": 800,
            "proportion": 0.7944588917783557
          },
          {
            "min": 800,
            "max": 1800,
            "proportion": 0.1080216043208642
          },
          {
            "min": 1800,
            "max": 0,
            "proportion": 0.09751950390077999
          }
        ],
        "category": "FAST"
      },
      "FIRST_CONTENTFUL_PAINT_MS": {
        "percentile": 1521,
        "distributions": [
          {
            "min": 0,
            "max": 1800,
            "proportion": 0.8027352460234876
          },
          {
            "min": 1800,
            "max": 3000,
            "proportion": 0.13349189832020217
          },
          {
            "min": 3000,
            "max": 0,
            "proportion": 0.0637728556563103
          }
        ],
        "category": "FAST"
      },
      "LARGEST_CONTENTFUL_PAINT_MS": {
        "percentile": 1643,
        "distributions": [
          {
            "min": 0,
            "max": 2500,
            "proportion": 0.8943399044205497
          },
          {
            "min": 2500,
            "max": 4000,
            "proportion": 0.08624551971326162
          },
          {
            "min": 4000,
            "max": 0,
            "proportion": 0.019414575866188773
          }
        ],
        "category": "FAST"
      }
    },
    "overall_category": "FAST",
    "initial_url": "https://aligoren.com/",
    "origin_fallback": false
  },
  "lighthouseResult": {
    "requestedUrl": "https://aligoren.com/",
    "finalUrl": "https://aligoren.com/",
    "lighthouseVersion": "9.3.0",
    "userAgent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) HeadlessChrome/98.0.4758.102 Safari/537.36",
    "fetchTime": "2022-05-20T17:10:52.157Z",
    "environment": {
      "networkUserAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4695.0 Safari/537.36 Chrome-Lighthouse",
      "hostUserAgent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) HeadlessChrome/98.0.4758.102 Safari/537.36",
      "benchmarkIndex": 772
    },
    "runWarnings": []
  }
}
```