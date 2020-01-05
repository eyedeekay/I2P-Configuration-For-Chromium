function platformCallback(platformInfo) {
  if (platformInfo.PlatformOs == "android") {
    console.log("android detected");
    return true;
  } else {
    console.log("desktop detected");
    return false;
  }
}

function isDroid() {
  return chrome.runtime.getPlatformInfo(platformCallback);
}
