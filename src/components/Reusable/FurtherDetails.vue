<template>
  <div id="further_detail">
    <!-- Close Button -->
    <div id="close_btn" @click="closeProjectDetail">X</div>
    <div id="details">
      <!-- Title -->
      <section id="project_title" v-if="detail.title !== ''">
        <h1>{{ detail.title }}</h1>
      </section>
      <!-- Description -->
      <section id="project_desc" v-if="detail.desc !== ''">
        <p>{{ detail.desc }}
          <a v-if="detail.organisation.name !== '' && detail.organisation.extra_links !== ''"
            :href="detail.organisation.extra_links"
            target="_blank">{{ detail.organisation.name }}
          </a>
        </p>
      </section>
      <!-- Image -->
      <section class="image_container" v-if="detail.images !== ''">
        <img :src="require(`@/assets/imgs/${detail.images}.png`)" />
      </section>

      <section id="links_group" >
        <!-- Videos -->
        <section class="detail" v-if="detail.videos.length !== 0">
          <div class="videos_container" v-for="(video, videoKey) in detail.videos"
            :key="videoKey">
            <h3 class="headers">{{ video.name }}</h3>
            <div class="icon_container">
              <a :href="video.link" target="_blank" alt="demo_video">
                <img src="@/assets/logos/youtube.png" />
              </a>
            </div>
          </div>
        </section>
        <!-- Project Log -->
        <section class="detail" v-if="detail.project_log !== ''">
          <h3 class="headers">Project Log</h3>
          <div class="icon_container">
            <a :href="detail.project_log" target="_blank" alt="report_link">
              <img src="@/assets/logos/project_log.png" />
            </a>
          </div>
        </section>
        <!-- Reports -->
        <section class="detail" v-if="detail.reports.length !== 0">
          <div class="report_container" v-for="(report, reportKey) in detail.reports"
            :key="reportKey">
            <h3 class="headers">{{ report.name }}</h3>
            <div class="icon_container">
              <a :href="report.link" target="_blank" alt="report_link">
                <img src="@/assets/logos/report.png" />
              </a>
            </div>
          </div>
        </section>
        <!-- Github -->
        <section class="detail" v-if="detail.github !== ''">
          <h3 class="headers">Github URL</h3>
          <div class="icon_container">
            <a :href="detail.github" target="_blank" alt="github_url">
              <img src="@/assets/logos/github.png" />
            </a>
          </div>
        </section>
        <!-- Live URL -->
        <section class="detail" v-if="detail.live_url !== ''">
          <h3 class="headers">{{ detail.live_url.name }}</h3>
          <div class="icon_container">
            <a :href="detail.live_url.link" target="_blank" alt="live_url">
              <img src="@/assets/logos/url.png" />
            </a>
          </div>
        </section>
      </section>
    </div>
  </div>
</template>

<script>
import { reactive } from 'vue';

export default {
  name: 'FurtherDetail',
  props: {
    emitter: {
      type: Object,
    },
    furtherDetail: {
      type: Object,
    },
  },
  setup(props) {
    const detail = reactive({
      title: props.furtherDetail.title,
      desc: props.furtherDetail.description,
      role: props.furtherDetail.role,
      images: props.furtherDetail.images.thumbnail,
      videos: props.furtherDetail.extra_links.videos,
      presentations: props.furtherDetail.extra_links.presentations,
      project_log: props.furtherDetail.extra_links.project_log,
      reports: props.furtherDetail.extra_links.reports,
      github: props.furtherDetail.github,
      live_url: props.furtherDetail.live_url,
      organisation: props.furtherDetail.organisation,
    });

    function closeProjectDetail() {
      props.emitter.emit('project_card', null);
    }

    return {
      closeProjectDetail,
      detail,
    };
  },
};
</script>

<style lang="scss" scoped>
#further_detail {
  display: block;
  position: relative;
  top: 15%;
  width: 80%;
  padding: 20px;
  text-align: right;

  #close_btn {
    display: inline-block;
    position: relative;
    top: 15%;
    padding: 20px;
    border-radius: 50%;
    border: 4px dashed black;
    cursor: pointer;
    transition: .2s all;

    &:hover {
      border: 4px dashed white;
      color: white;
    }
  }

  #details {
    position: relative;
    width: 100%;
    margin: 10px;
    padding: 20px;
    color: rgb(0, 0, 0);
    border-radius: 5px;
    background-color: rgba(87, 87, 87, 0.233);
    text-align: left;

    .image_container {
      position: relative;
      width: 100%;
      margin: 0;

      img {
        width: 100%;
        border-radius: 10px;
        box-shadow: 0px 0px 10px black;
      }
    }

    section {
      margin: 20px;
      text-align: center;
    }

    #project_title {
      text-align: center;
      background-color: rgba(255, 250, 250, 0.75);
      padding: 5px;
      border-radius: 5px;
    }
    #project_desc {
      text-align: left;
    }
    #links_group {
      position: relative;
      display: flex;
      flex-direction: row;
      align-items: center;
      justify-content: center;

      .detail {
        position: relative;
        width: 50%;

        .headers {
          margin: 5px;
        }
      }
    }

    .icon_container {
      position: relative;
      width: 100px;
      cursor: pointer;
      left: 50%;
      transform: translateX(-50%);

      img {
        width: 100%;
        transition: 0.2s filter;
      }

      &:hover {
        img {
          filter: invert(100%);
        }
      }
    }
  }
}
</style>
