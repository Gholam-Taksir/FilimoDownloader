<div align="center">

# 🎬 FilimoDownloader

**دانلودر فیلم و سریال از فیلیمو**

[![Go Version](https://img.shields.io/badge/Go-1.24+-00ADD8?style=flat&logo=go)](https://golang.org)
[![Platform](https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-lightgrey?style=flat)](https://github.com/Gholam-Taksir/FilimoDownloader)
[![License](https://img.shields.io/badge/License-MIT-green?style=flat)](LICENSE)
[![Version](https://img.shields.io/badge/Version-v2.0.0-red?style=flat)](https://github.com/Gholam-Taksir/FilimoDownloader/releases)

<br>

> ابزاری برای دانلود فیلم و سریال از [فیلیمو](https://www.filimo.com) با قابلیت انتخاب کیفیت، چند زبان صوتی، و زیرنویس

</div>

---

## ✨ قابلیت‌ها

- 🎬 دانلود فیلم با چند کیفیت (360p / 480p / 720p / 1080p)
- 📺 دانلود قسمت سریال
- 🔊 دانلود چند زبان صوتی همزمان (مثلاً فارسی + انگلیسی با هم)
- 💬 دانلود چند زیرنویس همزمان و embed در فایل نهایی
- ⏸ ادامه دانلود قطع‌شده (Resume)
- 📁 حذف خودکار فایل‌های موقت بعد از build
- 📋 تاریخچه دانلودها
- ⚡ سرعت بالا با buffer بهینه‌شده و retry خودکار

---

## 📋 پیش‌نیازها

قبل از هر چیز باید دو ابزار زیر نصب باشند:

---

### ۱. نصب Go

Go زبان برنامه‌نویسیه که برنامه باهاش نوشته شده — برای build کردن لازمه.

**ویندوز:**
1. برو به **[go.dev/dl](https://go.dev/dl/)**
2. فایل `.msi` مناسب ویندوزت رو دانلود کن (مثلاً `go1.24.3.windows-amd64.msi`)
3. فایل رو اجرا کن و مراحل نصب رو طی کن
4. بعد از نصب، یه CMD یا PowerShell جدید باز کن و تایپ کن:
```
go version
```
اگه چیزی مثل `go version go1.24.3 windows/amd64` نشون داد، نصب موفق بوده ✓

**لینوکس (Ubuntu/Debian):**
```bash
sudo apt update
sudo apt install golang-go
```
یا از سایت رسمی آخرین نسخه رو دانلود کن:
```bash
wget https://go.dev/dl/go1.24.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.24.3.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

**مک:**
```bash
brew install go
```
یا از [go.dev/dl](https://go.dev/dl/) فایل `.pkg` رو دانلود کن.

---

### ۲. نصب FFmpeg

FFmpeg برای ترکیب ویدیو، صدا و زیرنویس استفاده میشه — بدون این، فایل نهایی ساخته نمیشه.

**ویندوز:**
1. برو به **[ffmpeg.org/download.html](https://ffmpeg.org/download.html)**
2. روی آیکون ویندوز کلیک کن
3. از سایت **gyan.dev** گزینه `ffmpeg-release-essentials.zip` رو دانلود کن
4. zip رو extract کن — مثلاً در `C:\ffmpeg`
5. حالا باید FFmpeg رو به PATH اضافه کنی:
   - روی **This PC** راست‌کلیک کن → **Properties**
   - **Advanced system settings** → **Environment Variables**
   - در بخش **System variables** روی **Path** دبل‌کلیک کن
   - **New** رو بزن و مسیر پوشه `bin` رو بنویس، مثلاً: `C:\ffmpeg\bin`
   - OK → OK → OK
6. یه CMD جدید باز کن و تایپ کن:
```
ffmpeg -version
```
اگه اطلاعات نسخه نشون داد، نصب موفق بوده ✓

**لینوکس (Ubuntu/Debian):**
```bash
sudo apt update
sudo apt install ffmpeg
```

**لینوکس (Fedora/RHEL):**
```bash
sudo dnf install ffmpeg
```

**مک:**
```bash
brew install ffmpeg
```

---

## 🚀 نصب برنامه

**مرحله ۱ — دانلود سورس کد**

از بخش **[Releases](https://github.com/Gholam-Taksir/FilimoDownloader/releases)** آخرین نسخه رو دانلود کن، یا:

```bash
git clone https://github.com/Gholam-Taksir/FilimoDownloader.git
cd FilimoDownloader
```

**مرحله ۲ — Build کردن**

دانلود FilimoDownloader.exe از قسمت realese قرار دادن داخل پوشه سورس :
---

## 🔑 گرفتن توکن از فیلیمو

برای استفاده از برنامه باید اشتراک فعال فیلیمو داشته باشی و توکن احراز هویت رو بگیری:

1. مرورگرت رو باز کن و برو به [filimo.com](https://www.filimo.com)
2. با حساب کاربری‌ات **وارد شو**
3. کلید **F12** رو بزن تا DevTools باز بشه
4. روی تب **Application** کلیک کن
5. از منوی سمت چپ **Cookies** رو باز کن و `https://www.filimo.com` رو انتخاب کن
6. دنبال **`AuthV1`** بگرد
7. مقدار ستون **Value** رو کامل کپی کن

> 📌 این توکن در پوشه `data/ProfileDoc` ذخیره میشه و دفعات بعد نیازی به وارد کردن مجدد نیست.

---

## 📖 استفاده

```bash
# اجرای ساده — همه چیز رو میپرسه
FilimoDownloader.exe

# با شناسه مستقیم
FilimoDownloader.exe -i 12345

# با URL کامل
FilimoDownloader.exe -i https://www.filimo.com/m/12345

# با توکن مستقیم (بدون ذخیره)
FilimoDownloader.exe -t YOUR_TOKEN

# نمایش تاریخچه دانلودها
FilimoDownloader.exe --history

# Build دستی یه پوشه دانلود‌شده قبلی
FilimoDownloader.exe -b Downloads/MovieName

# تعیین پوشه خروجی
FilimoDownloader.exe -i 12345 -o "D:\Movies\MyMovie"
```

### همه گزینه‌ها

| فلگ | توضیح |
|-----|-------|
| `-i`, `--id` | شناسه یا آدرس کامل محتوا |
| `-t`, `--token` | توکن احراز هویت |
| `-b`, `--build` | Build پوشه دانلود‌شده |
| `-o`, `--output` | پوشه یا نام فایل خروجی |
| `-v`, `--version` | نمایش نسخه برنامه |
| `-h`, `--help` | نمایش راهنما |
| `--history` | نمایش تاریخچه دانلودها |

---

## 🎯 مراحل دانلود

وقتی برنامه رو اجرا میکنی، این مراحل رو طی میکنی:

```
1. نوع محتوا رو انتخاب کن:
   1) Movie
   2) Series - Single episode

2. شناسه یا URL رو وارد کن

3. کیفیت ویدیو رو انتخاب کن:
   1) 360p  2) 720p  3) 1080p

4. زبان صدا رو انتخاب کن (میتونی چند تا باهم):
   1) fa    2) en
   → مثال: 1,2 (هر دو زبان)

5. زیرنویس رو انتخاب کن (میتونی چند تا باهم):
   1) fa    2) en
   → مثال: 1,2 (هر دو زیرنویس)

6. دانلود شروع میشه...

7. بعد از دانلود، فرمت خروجی رو انتخاب کن:
   1) MP4    2) MKV

8. فایل نهایی ساخته میشه ✓
```

---

## 📁 ساختار پوشه بعد از دانلود

```
FilimoDownloader/
├── FilimoDownloader.exe
├── data/
│   ├── config.json              ← تنظیمات برنامه
│   ├── ProfileDoc               ← توکن ذخیره‌شده
│   └── download_history.json    ← تاریخچه دانلودها
└── Downloads/
    └── اسم_فیلم/
        ├── اسم_فیلم.mp4        ← فایل نهایی
        └── subtitle/
            ├── fa/
            │   └── subtitle.srt  ← زیرنویس فارسی
            └── en/
                └── subtitle.srt  ← زیرنویس انگلیسی
```

> 📌 پوشه‌های `video/` و `audio/` بعد از build به صورت خودکار حذف میشن

---

## ⚙️ فایل تنظیمات

فایل `data/config.json` بعد از اولین اجرا ساخته میشه:

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

| کلید | توضیح |
|------|-------|
| `default_quality` | کیفیت پیش‌فرض |
| `default_format` | فرمت خروجی (mp4/mkv) |
| `download_path` | مسیر پوشه دانلود |
| `auto_open_folder` | باز کردن پوشه بعد از دانلود |
| `show_info_before_dl` | نمایش اطلاعات فیلم قبل از دانلود |

---

## 🐛 مشکلات رایج

**❌ توکن نامعتبر**
توکن منقضی شده. دوباره از مرورگر کپی کن — فایل `data/ProfileDoc` رو حذف کن و برنامه رو اجرا کن.

**❌ FFmpeg پیدا نشد**
FFmpeg رو نصب کن و مطمئن شو به PATH اضافه شده. در CMD تایپ کن: `ffmpeg -version`

**❌ دانلود قطع شد**
نگران نباش — برنامه رو دوباره اجرا کن، از همون جایی که متوقف شده ادامه میده.

**❌ زیرنویس دانلود نشد**
برخی محتواها زیرنویس ندارن. از صفحه فیلیمو بررسی کن.

**❌ خطا در build**
مطمئن شو Go نصب هست: `go version`

---

## 🤝 مشارکت

Issue و Pull Request خوشامد است!

1. پروژه رو Fork کن
2. Branch جدید بساز: `git checkout -b feature/AmazingFeature`
3. تغییرات رو Commit کن: `git commit -m 'Add AmazingFeature'`
4. Push کن: `git push origin feature/AmazingFeature`
5. Pull Request بفرست

---

## ⭐ حمایت

اگه این ابزار برات مفید بود، یه **ستاره** بزن — خیلی انگیزه میده! 🙏

---

<div align="center">

ساخته‌شده با ❤️ توسط [Gholam-Taksir](https://github.com/Gholam-Taksir)

</div>
