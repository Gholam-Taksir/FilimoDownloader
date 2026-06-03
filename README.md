<div align="center">

# 🎬 FilimoDownloader

**دانلودر فیلم و سریال از فیلیمو**

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go)](https://golang.org)
[![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-lightgrey?style=flat)](https://github.com)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat)](LICENSE)
[![Version](https://img.shields.io/badge/Version-v2.0.0-red?style=flat)](https://github.com)

</div>

---

## ✨ قابلیت‌ها

- 🎬 دانلود فیلم با چند کیفیت (360p / 480p / 720p / 1080p)
- 📺 دانلود قسمت سریال
- 🔊 دانلود چند زبان صوتی همزمان (فارسی + انگلیسی)
- 💬 دانلود و embed زیرنویس (چند زبان همزمان)
- ⏸ ادامه دانلود قطع‌شده (Resume)
- 📁 حذف خودکار فایل‌های موقت بعد از build
- 📋 تاریخچه دانلودها
- ⚡ سرعت بالا با buffer بهینه‌شده

---

## 📋 پیش‌نیازها

| ابزار | لینک دانلود |
|-------|-------------|
| **FFmpeg** | [ffmpeg.org/download](https://ffmpeg.org/download.html) |
| **Go 1.24+** *(فقط برای build)* | [go.dev/dl](https://go.dev/dl/) |

> ⚠️ FFmpeg باید در PATH سیستم باشد

---

## 🚀 نصب و راه‌اندازی

### روش ۱ — دانلود مستقیم (توصیه‌شده)

از بخش **[Releases](../../releases)** آخرین نسخه را دانلود کنید:

```
FilimoDownloader-windows.zip   ← ویندوز
FilimoDownloader-linux.tar.gz  ← لینوکس
FilimoDownloader-mac.tar.gz    ← مک
```

### روش ۲ — Build از سورس

```bash
git clone https://github.com/YOUR_USERNAME/FilimoDownloader.git
cd FilimoDownloader
go build -ldflags="-X main.isProduction=true" -o FilimoDownloader.exe ./cmd/FilimoDownloader
```

یا در ویندوز:
```
build.bat
```

---

## 🔑 گرفتن توکن

1. وارد [filimo.com](https://www.filimo.com) شوید
2. **F12** بزنید (DevTools)
3. به تب **Application** بروید
4. از منوی **Cookies** → `filimo.com` را انتخاب کنید
5. مقدار **`AuthV1`** را کپی کنید

---

## 📖 استفاده

```bash
# اجرای ساده
FilimoDownloader.exe

# با ID مستقیم
FilimoDownloader.exe -i 12345

# با URL کامل
FilimoDownloader.exe -i https://www.filimo.com/m/12345

# با توکن مستقیم
FilimoDownloader.exe -t YOUR_TOKEN

# نمایش تاریخچه
FilimoDownloader.exe --history

# Build دستی یه پوشه دانلود‌شده
FilimoDownloader.exe -b Downloads/MovieName
```

### گزینه‌ها

| فلگ | توضیح |
|-----|-------|
| `-i`, `--id` | شناسه یا آدرس محتوا |
| `-t`, `--token` | توکن احراز هویت |
| `-b`, `--build` | Build پوشه دانلود‌شده |
| `-o`, `--output` | پوشه خروجی |
| `-v`, `--version` | نمایش نسخه |
| `-h`, `--help` | راهنما |

---

## 📁 ساختار پوشه

```
FilimoDownloader/
├── FilimoDownloader.exe
├── data/
│   ├── config.json           ← تنظیمات
│   ├── ProfileDoc            ← توکن ذخیره‌شده
│   └── download_history.json ← تاریخچه
└── Downloads/
    └── MovieName/
        ├── MovieName.mp4     ← فایل نهایی
        └── subtitle/
            ├── fa/subtitle.srt
            └── en/subtitle.srt
```

---

## ⚙️ تنظیمات (data/config.json)

```json
{
  "default_quality": "720p",
  "default_format": "mp4",
  "download_path": "Downloads",
  "max_threads": 2,
  "auto_open_folder": true,
  "show_info_before_dl": true
}
```

---

## 🐛 مشکلات رایج

**توکن نامعتبر:**
> توکن را حذف کرده و دوباره از مرورگر کپی کنید

**FFmpeg پیدا نشد:**
> FFmpeg را نصب کرده و به PATH اضافه کنید

**دانلود قطع شد:**
> برنامه را دوباره اجرا کنید — از همان جایی که متوقف شده ادامه می‌دهد

**زیرنویس دانلود نشد:**
> برخی محتواها زیرنویس ندارند. از صفحه فیلیمو بررسی کنید

---

## 🤝 مشارکت

Pull Request و Issue خوشامد است!

---

## ⭐ حمایت

اگه این ابزار برات مفید بود، یه **ستاره** بزن — خیلی کمک میکنه! 🙏

---

<div align="center">
ساخته‌شده با ❤️ توسط <a href="https://github.com/GholamTaksir">GholamTaksir</a>
</div>
