<template>
  <div id="nav">
    <ul id="els">
      <li class="nav_el" id="home">Home</li>
      <li class="nav_el" id="portfolio">Portfolio</li>
      <li class="nav_el" id="blog">Blog</li>
      <li class="nav_el" id="contact">Resume</li>
    </ul>
  </div>
</template>

<script>

export default {
  name: 'Navigation',
  setup() {
  },
};
</script>

<style lang="scss" scoped>
#nav {
  position: fixed;
  top: 0;
  width: 100%;
  background-color: black;
  border-bottom-left-radius: 25%;
  border-bottom-right-radius: 25%;
  z-index: 1;

  #els {
    list-style: none;
    padding: 20px;

    .nav_el {
      display: inline-block;
      margin: 20px;
      color: rgb(218, 218, 218);
      font-size: 1.25rem;
      cursor: pointer;
    }
  }
}
</style>
