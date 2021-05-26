<template>
  <div id="footer">
    <footer>Made with <a href="https://v3.vuejs.org/" target="_blank">Vue.js</a></footer>
  </div>
</template>

<script>
export default {
  name: 'Footer',
};
</script>

<style lang="scss" scoped>
#footer {
  position: absolute;
  bottom: 0;
  margin: 0 auto;
  padding: 10px;
  background-color: rgb(148, 53, 53);
  color: rgb(183, 183, 233);
  width: 100%;

  a {
    text-decoration: none;
    color: rgb(178, 178, 245);
  }
}

/* Responsive Text */
@media screen and (max-width: 770px) {
  #footer {
    font-size: 15px;
  }
}
@media screen and (max-width: 600px) {
  #footer {
    font-size: 12px;
  }
}
@media screen and (max-width: 460px) {
  #footer {
    font-size: 9px;
  }
}
</style>
