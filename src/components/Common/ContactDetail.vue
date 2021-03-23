<template>
  <div id="contact">
    <div id="outer_wrapper">
      <div class="img_container" id="mail">
        <img src="@/assets/mail_original.png" />
      </div>
      <div class="img_container" id="youtube">
        <img src="@/assets/youtube.png" />
      </div>
      <div class="img_container" id="linkedin">
        <img src="@/assets/linkedin.png" />
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ContactDetail',
  setup() {

  },
};
</script>

<style lang="scss" scoped>
#contact {
  position: fixed;
  top: 0;
  left: 0;
  width: 20%;
  height: 100vh;
  background-color: rgb(226, 91, 91);
  z-index: -1;
  #outer_wrapper {
    position: absolute;
    top: 50%;
    transform: translateY(-50%);
    .img_container {
      width: 20%;
      margin: 40px;

      img {
        width: 100%;
      }
    }
  }
}
</style>
