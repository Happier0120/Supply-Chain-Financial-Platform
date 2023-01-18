<template>
  <div class="order">
    <router-link to="/order/orderList"></router-link>
    <router-link to="/order/createOrder"></router-link>
    <router-link to="/order/editOrder"></router-link>
    <router-link to="/order/queryOrder"></router-link>
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  data() {
    return {};
  },
};
</script>
