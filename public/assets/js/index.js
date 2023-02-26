if (localStorage.getItem("token") !== null) {
  location.href = "/health";
}

import { createApp, ref } from "vue";

createApp({
  data() {
    return {
      message: "Hello Vue!",
    };
  },
}).mount("#app");
