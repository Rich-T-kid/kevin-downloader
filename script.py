import yt_dlp

# Simple download
url = ["https://www.youtube.com/watch?v=dQw4w9WgXcQ"]

with open("links.txt", "r") as f:
    url = f.read().splitlines()
def download():
    count = 0
    ydl_opts = {
        "outtmpl": "mussie-downloads/%(title)s.%(ext)s",
        "format": "bestaudio/best",
        "cookiefile": "cookies.txt",
        "ignoreerrors": True,  # Skip unavailable videos
        "no_warnings": False,
    }
    
    with yt_dlp.YoutubeDL(ydl_opts) as ydl:
        for video_url in url:
            try:
                print(f"\n[{count + 1}/{len(url)}] Downloading: {video_url}")
                ydl.download([video_url])
                count += 1
            except yt_dlp.utils.DownloadError as e:
                print(f"Skipping video due to error: {e}")
                continue
            except Exception as e:
                print(f"Skipping video due to unexpected error: {e}")
                continue
    
    print(f"\n✓ Download complete! Successfully processed {count}/{len(url)} videos")
download()